package conf

import "context"

type ILogger interface {
	Info(v ...any)
	Infoln(v ...any)
	Infof(format string, v ...any)

	Warn(v ...any)
	Warnln(v ...any)
	Warnf(format string, v ...any)

	Error(v ...any)
	Errorln(v ...any)
	Errorf(format string, v ...any)

	InfoContext(ctx context.Context, v ...any)
	InfoContextf(ctx context.Context, format string, v ...any)

	WarnContext(ctx context.Context, v ...any)
	WarnContextf(ctx context.Context, format string, v ...any)

	ErrorContext(ctx context.Context, v ...any)
	ErrorContextf(ctx context.Context, format string, v ...any)
}

var i ILogger

func Internal() ILogger {
	if i == nil {
		checkLogger()
	}
	return i
}
