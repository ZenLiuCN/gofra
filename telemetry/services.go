package telemetry

import (
	"context"
	"fmt"
	"github.com/ZenLiuCN/gofra/conf"
	rt "go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"time"
)

const (
	Version = "0.0.1"
)
const ContextKey = "$telemetry"

func FromContext(ctx context.Context) (r Telemetry) {
	if ctx == nil {
		return nil
	}
	r, ok := ctx.Value(ContextKey).(Telemetry)
	if !ok {
		return nil
	}
	return r
}
func ByContext(ctx context.Context, opts ...trace.SpanStartOption) (r Telemetry) {
	if ctx == nil {
		return nil
	}
	r, ok := ctx.Value(ContextKey).(Telemetry)
	if !ok {
		return NewTelemetry("", opts...)
	}
	return r
}

type Telemetry interface {
	Scope() string
	HandleError(err error)

	/*
		HandleRecover returns the recover value and if there is failure
		A sample usage like:
			defer func() {
				if r, ok := tel.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
	*/
	HandleRecover(rec any) (failure any, failed bool)

	//StartSpan if ctx is nil then the returns are nil
	StartSpan(name string, ctx context.Context, attrs ...attribute.KeyValue) (context.Context, trace.Span)

	SetContext(ctx context.Context) context.Context
}
type telemetry struct {
	spanStartOption []trace.SpanStartOption
	scope           string
	propagator      propagation.TextMapPropagator
	meter           metric.Meter
	tracer          trace.Tracer
}

func (t *telemetry) StartSpan(name string, ctx context.Context, attrs ...attribute.KeyValue) (cx context.Context, sp trace.Span) {
	if ctx == nil {
		return ctx, nil
	}
	cx, sp = t.tracer.Start(ctx, name, t.spanStartOption...)
	if len(attrs) > 0 {
		sp.SetAttributes(attrs...)
	}
	return
}
func (t *telemetry) Scope() string {
	return t.scope
}
func (t *telemetry) HandleError(err error) {
	if err != nil {
		otel.Handle(err)
	}
}
func (t *telemetry) HandleRecover(rec any) (any, bool) {
	switch r := rec.(type) {
	case nil:
		return r, false
	case error:
		t.HandleError(r)
		return r, true
	default:
		t.HandleError(fmt.Errorf("%#+v", r))
		return r, true
	}
}

func (t *telemetry) SetContext(ctx context.Context) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, ContextKey, t)
}
func NewTelemetry(scope string, opts ...trace.SpanStartOption) Telemetry {
	return &telemetry{
		scope:           scope,
		spanStartOption: opts,
		propagator:      otel.GetTextMapPropagator(),
		meter:           otel.GetMeterProvider().Meter(scope, metric.WithInstrumentationVersion(Version)),
		tracer:          otel.GetTracerProvider().Tracer(scope, trace.WithInstrumentationVersion(Version)),
	}
}

type ServiceFunc func(ctx context.Context)

func Instrument(name string, service ServiceFunc) ServiceFunc {
	return func(ctx context.Context) {
		if tel := ByContext(ctx); tel != nil {
			cx, span := tel.StartSpan(name, ctx)
			defer func() {
				if r, ok := tel.HandleRecover(recover()); ok {
					panic(r)
				}
			}()
			defer span.End()
			service(cx)
		} else {
			service(ctx)
		}

	}
}

func RuntimeInstrument(c conf.Config) {
	if c == nil {
		Handle(rt.Start(rt.WithMinimumReadMemStatsInterval(time.Second)))
	} else {
		Handle(rt.Start(rt.WithMinimumReadMemStatsInterval(c.GetTimeDurationInfiniteNotAllowed("telemetry.runtime.interval", time.Second))))
	}

}
func Handle(err error) {
	if err != nil {
		otel.Handle(err)
	}
}
