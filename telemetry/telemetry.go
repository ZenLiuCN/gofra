package telemetry

import (
	"context"
	"errors"
	"github.com/ZenLiuCN/goinfra/conf"
	"github.com/ZenLiuCN/goinfra/telemetry/otlp"
	"github.com/ZenLiuCN/goinfra/telemetry/prometheus"

	"go.opentelemetry.io/otel"
)

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
	tracerProvider, err := otlp.NewTraceProvider(conf)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFunc = append(shutdownFunc, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Set up meter provider.
	meterProvider, err := prometheus.NewMeterProvider(conf)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFunc = append(shutdownFunc, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)
	return
}
