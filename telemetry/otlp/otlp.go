package otlp

import (
	"github.com/ZenLiuCN/goinfra/conf"
	otlp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"time"
)

func NewTraceProvider(cfg conf.Config) (*trace.TracerProvider, error) {
	var opt []otlp.Option
	opt = append(opt, otlp.WithEndpointURL(conf.Required("telemetry.oltp.endpoint", cfg, cfg.GetString)))
	conf.Exists("telemetry.otlp.compress", cfg, cfg.GetString, func(d string) {
		opt = append(opt, otlp.WithCompressor(d))
	})
	traceExporter := otlp.NewUnstarted(opt...)
	var opts []trace.TracerProviderOption
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
	traceProvider := trace.NewTracerProvider(opts...)
	return traceProvider, nil
}
