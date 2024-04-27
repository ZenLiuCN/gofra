package otlp

import (
	"context"
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/ZenLiuCN/gofra/telemetry/common"
	otlp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"log/slog"
	"time"
)

func NewTraceProvider(ctx context.Context, cfg conf.Config) (*trace.TracerProvider, error) {
	var opt []otlp.Option
	{
		opt = append(opt, otlp.WithEndpointURL(conf.Required("telemetry.oltp.endpoint", cfg, cfg.GetString)))
		if d := cfg.GetString("telemetry.otlp.compress", ""); d != "" {
			opt = append(opt, otlp.WithCompressor(d))
		}
		if cfg.GetBoolean("telemetry.otlp.insecure", false) {
			opt = append(opt, otlp.WithInsecure())
		}
		if d := cfg.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.reconnect"); d != time.Duration(0) {
			opt = append(opt, otlp.WithReconnectionPeriod(d))
		}
		if d := cfg.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.timeout"); d != time.Duration(0) {
			opt = append(opt, otlp.WithTimeout(d))
		}
		if d := cfg.GetObject("telemetry.otlp.retry"); d != nil {
			var retry otlp.RetryConfig
			retry.Enabled = cfg.GetBoolean("telemetry.otlp.retry.enable", true)
			retry.InitialInterval = cfg.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.retry.initialInterval", time.Second*5)
			retry.MaxInterval = cfg.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.retry.maxInterval", time.Second*30)
			retry.MaxElapsedTime = cfg.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.retry.maxElapsedTime", time.Minute)
			opt = append(opt, otlp.WithRetry(retry))
		}
		if d := cfg.GetObject("telemetry.otlp.headers"); d != nil {
			mp := d.GetStringMap("")
			if len(mp) != 0 {
				m := make(map[string]string, len(mp))
				for s, config := range mp {
					m[s] = config.GetString("")
				}
				opt = append(opt, otlp.WithHeaders(m))
			}
		}
	}
	traceExporter, err := otlp.New(ctx, opt...)
	if err != nil {
		return nil, err
	}
	var opts []trace.TracerProviderOption
	{
		//!! batch
		{
			var spanOpt []trace.BatchSpanProcessorOption
			conf.Exists("telemetry.otlp.trace.export.timeout", cfg, cfg.GetTimeDuration, func(d time.Duration) {
				spanOpt = append(spanOpt, trace.WithExportTimeout(d))
			})
			conf.Exists("telemetry.otlp.trace.export.size", cfg, cfg.GetInt32, func(d int32) {
				spanOpt = append(spanOpt, trace.WithMaxExportBatchSize(int(d)))
			})
			conf.Exists("telemetry.otlp.trace.batch.size", cfg, cfg.GetInt32, func(d int32) {
				spanOpt = append(spanOpt, trace.WithMaxExportBatchSize(int(d)))
			})
			conf.Exists("telemetry.otlp.trace.batch.timeout", cfg, cfg.GetTimeDuration, func(d time.Duration) {
				spanOpt = append(spanOpt, trace.WithBatchTimeout(d))
			})
			conf.Exists("telemetry.otlp.trace.queue.size", cfg, cfg.GetInt32, func(d int32) {
				spanOpt = append(spanOpt, trace.WithMaxQueueSize(int(d)))
			})
			conf.Exists("telemetry.otlp.trace.queue.blocking", cfg, cfg.GetBoolean, func(d bool) {
				if d {
					spanOpt = append(spanOpt, trace.WithBlocking())
				}
			})
			opts = append(opts, trace.WithBatcher(traceExporter, spanOpt...))
		}
		//!! resource
		{
			res, err := common.ParseResource(ctx, cfg)
			if err != nil {
				return nil, err
			}
			opts = append(opts, trace.WithResource(res))
		}
		//!! sampler
		{
			if sampler := cfg.GetObject("telemetry.otlp.trace.sample"); sampler != nil {
				switch sampler.GetString("name") {
				case "always":
					opts = append(opts, trace.WithSampler(trace.AlwaysSample()))
				case "never":
					opts = append(opts, trace.WithSampler(trace.NeverSample()))
				case "parent":
					var options []trace.ParentBasedSamplerOption
					var sam trace.Sampler
					switch sampler.GetString("based") {
					case "always":
						sam = trace.AlwaysSample()
					case "ratio":
						sam = trace.TraceIDRatioBased(sampler.GetFloat64("ratio", 1))
					case "never":
						sam = trace.NeverSample()
					default:
						slog.Error("telemetry.oltp.trace.sample.based not one of always|never|ratio, will use never as default",
							"based", sampler.GetString("based"),
						)
						sam = trace.NeverSample()
					}
					for _, s := range sampler.GetStringList("options") {
						switch s {
						case "withRemote":
							options = append(options, trace.WithRemoteParentSampled(sam))
						case "withoutRemote":
							options = append(options, trace.WithRemoteParentNotSampled(sam))
						case "withLocal":
							options = append(options, trace.WithLocalParentSampled(sam))
						case "withoutLocal":
							options = append(options, trace.WithLocalParentNotSampled(sam))
						default:
							slog.Warn("unknown options", "name", s)
						}
					}
					opts = append(opts, trace.WithSampler(trace.ParentBased(sam, options...)))
				default:
					slog.Error("telemetry.oltp.trace.sample.name not one of always|never|parent, will use system default",
						"name", sampler.GetString("name"),
					)
				}
			}
		}

	}
	traceProvider := trace.NewTracerProvider(opts...)
	return traceProvider, nil
}
