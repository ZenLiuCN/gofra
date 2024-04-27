package prometheus

import (
	"context"
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/ZenLiuCN/gofra/telemetry/common"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func NewMeterProvider(ctx context.Context, c conf.Config) (*metric.MeterProvider, error) {
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
		res, err := common.ParseResource(ctx, c)
		if err != nil {
			return nil, err
		}
		mpo = append(mpo, metric.WithResource(res))
	}
	meterProvider := metric.NewMeterProvider(mpo...)
	return meterProvider, nil
}
