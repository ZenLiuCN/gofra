package telemetry

import (
	"context"
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/ZenLiuCN/ote"
	"github.com/ZenLiuCN/ote/otlp"
	"github.com/ZenLiuCN/ote/resource"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"time"
)

func SetupTelemetry(ctx context.Context, c conf.Config) (s func(context.Context) error, err error) {
	cx := new(otlp.TraceConfig)
	cx.Endpoint = c.RequiredString("telemetry.otlp.endpoint")

	cx.Compress = c.GetString("telemetry.otlp.compress")
	c.ExistsBoolean("telemetry.otlp.insecure", func(b bool) {
		cx.Insecure.Valid = true
		cx.Insecure.Bool = b
	})
	cx.Reconnect = c.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.reconnect")
	cx.Timeout = c.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.timeout")
	if c.HasPath("telemetry.otlp.retry") {
		cc := c.GetObject("telemetry.otlp.retry")
		cx.Retry = new(otlptracegrpc.RetryConfig)
		cx.Retry.Enabled = true
		cx.Retry.InitialInterval = 5 * time.Second
		cx.Retry.MaxInterval = 30 * time.Second
		cx.Retry.MaxElapsedTime = time.Minute
		cc.ExistsDuration("initDelay", func(d time.Duration) { cx.Retry.InitialInterval = d })
		cc.ExistsDuration("maxInterval", func(d time.Duration) { cx.Retry.MaxInterval = d })
		cc.ExistsDuration("maxElapsed", func(d time.Duration) { cx.Retry.MaxElapsedTime = d })
	}

	cx.Headers = c.GetTextMap("telemetry.otlp.headers")
	cx.ExportTimeout = c.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.export.timeout")
	c.ExistsInt32("telemetry.otlp.export.batch.size", func(i int32) {
		cx.ExportBatchSize.Valid = true
		cx.ExportBatchSize.Int32 = i
	})

	cx.ExportBatchTimeout = c.GetTimeDurationInfiniteNotAllowed("telemetry.otlp.export.batch.timeout")
	c.ExistsInt32("telemetry.otlp.export.queue.size", func(i int32) {
		cx.QueueSize.Valid = true
		cx.QueueSize.Int32 = i
	})
	c.ExistsBoolean("telemetry.otlp.export.queue.blocking", func(i bool) {
		cx.QueueBlocking.Valid = true
		cx.QueueBlocking.Bool = i
	})
	if c.HasPath("telemetry.otlp.sampler") {
		cc := c.GetObject("telemetry.otlp.sampler")
		cx.Sampler = new(otlp.SamplerConfig)
		cx.Sampler.Name = cc.RequiredString("name")
		cx.Sampler.Based = cc.GetString("base")
		cc.ExistsFloat64("ratio", func(d float64) { cx.Sampler.Ratio.Valid = true; cx.Sampler.Ratio.Float64 = d })
		if cc.HasPath("options") {
			cx.Sampler.Options = cc.GetStringList("options")
		}
	}
	cx.Config = new(resource.Config)
	if c.HasPath("telemetry.resource") {
		cc := c.GetObject("telemetry.resource")
		cc.ExistsString("service", func(s string) {
			cx.Config.Service.Valid = true
			cx.Config.Service.String = s
		})
		cc.ExistsBoolean("container", func(b bool) {
			cx.Config.Container.Valid = true
			cx.Config.Container.Bool = b
		})
		cc.ExistsBoolean("host", func(b bool) {
			cx.Config.Host.Valid = true
			cx.Config.Host.Bool = b
			cx.Config.HostId.Valid = true
			cx.Config.HostId.Bool = b
		})
		cc.ExistsBoolean("env", func(b bool) {
			cx.Config.Env.Valid = true
			cx.Config.Env.Bool = b
		})
		cc.ExistsBoolean("process", func(b bool) {
			cx.Config.Process.Valid = true
			cx.Config.Process.Bool = b
		})
		cc.ExistsBoolean("sdk", func(b bool) {
			cx.Config.SDK.Valid = true
			cx.Config.SDK.Bool = b
		})
	}
	if d := c.GetTimeDurationInfiniteNotAllowed("telemetry.runtime", time.Second); d != 0 {
		defer func() {
			ote.RuntimeInstrument(d)
		}()
	}
	return ote.SetupTelemetry(ctx, cx)
}
