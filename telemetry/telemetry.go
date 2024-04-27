package telemetry

import (
	"context"
	"errors"
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/ZenLiuCN/gofra/telemetry/otlp"
	"github.com/ZenLiuCN/gofra/telemetry/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel"
)

var (
	have bool
)

func UsingTelemetry() bool {
	return have
}
func SetupTelemetry(ctx context.Context, conf conf.Config) (shutdown func(context.Context) error, err error) {
	var shutdownFunc []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFunc {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFunc = nil
		return err
	}
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}
	prop := NewPropagator(conf)
	otel.SetTextMapPropagator(prop)
	var tracerProvider *trace.TracerProvider
	if tracerProvider, err = otlp.NewTraceProvider(ctx, conf); err != nil {
		handleErr(err)
		return
	} else {
		shutdownFunc = append(shutdownFunc, tracerProvider.Shutdown)
		otel.SetTracerProvider(tracerProvider)
	}
	var meterProvider *metric.MeterProvider
	if meterProvider, err = prometheus.NewMeterProvider(ctx, conf); err != nil {
		handleErr(err)
		return
	} else {
		shutdownFunc = append(shutdownFunc, meterProvider.Shutdown)
		otel.SetMeterProvider(meterProvider)
	}
	have = true
	return
}
