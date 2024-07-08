//go:build glog || !slog

package conf

import (
	"context"
	"github.com/golang/glog"
)

type adaptor struct {
}

func (w adaptor) Info(v ...any) {
	glog.Info(v...)
}

func (w adaptor) Infoln(v ...any) {
	glog.Infoln(v...)
}

func (w adaptor) Infof(format string, v ...any) {
	glog.Infof(format, v...)
}

func (w adaptor) Warn(v ...any) {
	glog.Warning(v...)
}

func (w adaptor) Warnln(v ...any) {
	glog.Warningln(v...)
}

func (w adaptor) Warnf(format string, v ...any) {
	glog.Warningf(format, v...)
}

func (w adaptor) Error(v ...any) {
	glog.Error(v...)
}

func (w adaptor) Errorln(v ...any) {
	glog.Errorln(v...)
}

func (w adaptor) Errorf(format string, v ...any) {
	glog.Errorf(format, v...)
}

func (w adaptor) InfoContext(ctx context.Context, v ...any) {
	glog.InfoContext(ctx, v...)
}

func (w adaptor) InfoContextf(ctx context.Context, format string, v ...any) {
	glog.InfoContextf(ctx, format, v...)
}

func (w adaptor) WarnContext(ctx context.Context, v ...any) {
	glog.WarningContext(ctx, v...)
}

func (w adaptor) WarnContextf(ctx context.Context, format string, v ...any) {
	glog.WarningContextf(ctx, format, v...)
}

func (w adaptor) ErrorContext(ctx context.Context, v ...any) {
	glog.ErrorContext(ctx, v...)
}

func (w adaptor) ErrorContextf(ctx context.Context, format string, v ...any) {

	glog.ErrorContextf(ctx, format, v...)
}

func checkLogger() {
	i = adaptor{}
}
