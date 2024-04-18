package prometheus

import (
	"context"
	"errors"
	"github.com/ZenLiuCN/goinfra/conf"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"log/slog"
)

func NewMeterProvider(c conf.Config) (*metric.MeterProvider, error) {
	var opt []prometheus.Option
	{

	}
	metricExporter, err := prometheus.New(opt...)
	if err != nil {
		return nil, err
	}

	var mpo []metric.Option
	{
		mpo = append(mpo, metric.WithReader(metricExporter))
		var res *resource.Resource
		res, err = resource.New(
			context.Background(),
			resource.WithContainer(),
			resource.WithHost(),
			resource.WithHostID(),
			resource.WithFromEnv(),
			resource.WithOS(),
			resource.WithProcess(),
			resource.WithTelemetrySDK(),
		)
		if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
			slog.Error("telemetry resource error", err)
		} else if err != nil {
			return nil, err
		}
		mpo = append(mpo, metric.WithResource(res))
	}
	meterProvider := metric.NewMeterProvider(mpo...)
	return meterProvider, nil
}
