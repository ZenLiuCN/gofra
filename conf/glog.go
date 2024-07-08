//go:build glog || !slog

package conf

import (
	"context"
	"github.com/golang/glog"
)

type adaptor struct {
}

func (w adaptor) Info(v ...any) {
	glog.InfoDepth(1, v...)
}

func (w adaptor) Infof(format string, v ...any) {
	glog.InfoDepthf(1, format, v...)
}

func (w adaptor) Warn(v ...any) {
	glog.WarningDepth(1, v...)
}

func (w adaptor) Warnf(format string, v ...any) {
	glog.WarningDepthf(1, format, v...)
}

func (w adaptor) Error(v ...any) {
	glog.ErrorDepth(1, v...)
}

func (w adaptor) Errorf(format string, v ...any) {
	glog.ErrorDepthf(1, format, v...)
}

func (w adaptor) InfoContext(ctx context.Context, v ...any) {
	glog.InfoContextDepth(ctx, 1, v...)
}

func (w adaptor) InfoContextf(ctx context.Context, format string, v ...any) {
	glog.InfoContextDepthf(ctx, 1, format, v...)
}

func (w adaptor) WarnContext(ctx context.Context, v ...any) {
	glog.WarningContextDepth(ctx, 1, v...)
}

func (w adaptor) WarnContextf(ctx context.Context, format string, v ...any) {
	glog.WarningContextDepthf(ctx, 1, format, v...)
}

func (w adaptor) ErrorContext(ctx context.Context, v ...any) {
	glog.ErrorContextDepth(ctx, 1, v...)
}

func (w adaptor) ErrorContextf(ctx context.Context, format string, v ...any) {
	glog.ErrorContextDepthf(ctx, 1, format, v...)
}

func checkLogger() {
	i = adaptor{}
}
