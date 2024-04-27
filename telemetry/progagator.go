package telemetry

import (
	"github.com/ZenLiuCN/gofra/conf"
	"go.opentelemetry.io/otel/propagation"
)

func NewPropagator(conf conf.Config) propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
