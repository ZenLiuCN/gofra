package common

import (
	"context"
	"errors"
	"github.com/ZenLiuCN/goinfra/conf"
	"github.com/ZenLiuCN/goinfra/utils"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"log/slog"
)

func ParseResource(ctx context.Context, c conf.Config) (res *resource.Resource, err error) {
	var opts []resource.Option
	{
		opts = append(opts, resource.WithAttributes(semconv.ServiceName(c.GetString("telemetry.resource.service", utils.ExecutableName()))))
		if c.GetBoolean("telemetry.resource.container", true) {
			opts = append(opts, resource.WithContainer())
		}
		if c.GetBoolean("telemetry.resource.host", true) {
			opts = append(opts, resource.WithHost())
		}
		if c.GetBoolean("telemetry.resource.hostID", true) {
			opts = append(opts, resource.WithHostID())
		}
		if c.GetBoolean("telemetry.resource.env", true) {
			opts = append(opts, resource.WithFromEnv())
		}
		if c.GetBoolean("telemetry.resource.process", true) {
			opts = append(opts, resource.WithProcess())
		}
		if c.GetBoolean("telemetry.resource.sdk", true) {
			opts = append(opts, resource.WithTelemetrySDK())
		}
	}
	res, err = resource.New(ctx, opts...)
	if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
		slog.Error("telemetry resource error", err)
		return res, nil
	} else if err != nil {
		return nil, err
	}
	return
}
