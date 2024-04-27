package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/ZenLiuCN/gofra/utils"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"log/slog"
)

var (
	global *resource.Resource
	flags  = 0
)

func computeFlag(c conf.Config) (bool, int) {
	f := 0
	if c.GetBoolean("telemetry.resource.container", true) {
		f &= 1 << 0
	}
	if c.GetBoolean("telemetry.resource.host", true) {
		f &= 1 << 1
	}
	if c.GetBoolean("telemetry.resource.hostID", true) {
		f &= 1 << 2
	}
	if c.GetBoolean("telemetry.resource.env", true) {
		f &= 1 << 3
	}
	if c.GetBoolean("telemetry.resource.process", true) {
		f &= 1 << 4
	}
	if c.GetBoolean("telemetry.resource.sdk", true) {
		f &= 1 << 5
	}
	same := f == flags
	return same, f
}
func ParseResource(ctx context.Context, c conf.Config) (res *resource.Resource, err error) {
	f := 0
	if global != nil {
		var same bool
		same, f = computeFlag(c)
		if same {
			res = global
			return
		}
	}
	var compute = f == 0
	var opts []resource.Option
	{

		opts = append(opts, resource.WithAttributes(semconv.ServiceName(c.GetString("telemetry.resource.service", utils.ExecutableName()))))
		if c.GetBoolean("telemetry.resource.container", true) {
			if compute {
				f &= 1 << 0
			}
			opts = append(opts, resource.WithContainer())
		}
		if c.GetBoolean("telemetry.resource.host", true) {
			if compute {
				f &= 1 << 1
			}
			opts = append(opts, resource.WithHost())
		}
		if c.GetBoolean("telemetry.resource.hostID", true) {
			if compute {
				f &= 1 << 2
			}
			opts = append(opts, resource.WithHostID())
		}
		if c.GetBoolean("telemetry.resource.env", true) {
			if compute {
				f &= 1 << 3
			}
			opts = append(opts, resource.WithFromEnv())
		}
		if c.GetBoolean("telemetry.resource.process", true) {
			if compute {
				f &= 1 << 4
			}
			opts = append(opts, resource.WithProcess())
		}
		if c.GetBoolean("telemetry.resource.sdk", true) {
			if compute {
				f &= 1 << 5
			}
			opts = append(opts, resource.WithTelemetrySDK())
		}
	}
	res, err = resource.New(ctx, opts...)
	if global == nil && compute {
		global = res
		flags = f
	}
	if errors.Is(err, resource.ErrPartialResource) || errors.Is(err, resource.ErrSchemaURLConflict) {
		slog.Error("telemetry resource error", "error", err)
		return res, nil
	} else if err != nil {
		slog.Error("telemetry resource fatal error", "error", err, "kind", fmt.Sprintf("%T", err), "res", res)
		return res, nil
	}
	return
}
