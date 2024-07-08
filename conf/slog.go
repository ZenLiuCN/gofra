//go:build slog && !glog

package conf

import (
	"context"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"log/slog"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	handler *RotateFileHandler
)

type RotateFileHandler struct {
	path    string
	pattern string
	file    *os.File
	limit   int64
	lock    sync.Mutex
}

func (s *RotateFileHandler) Write(p []byte) (n int, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, _ = os.Stdout.Write(p)
	n, err = s.file.Write(p)
	if err == nil {
		if si, er := s.file.Stat(); er == nil && si.Size() > s.limit {
			_ = s.file.Close()
			_ = os.Rename(s.path, strings.ReplaceAll(s.path, ".log", "."+time.Now().Format(s.pattern)+".log"))
			s.file, _ = os.OpenFile(s.path, os.O_APPEND|os.O_CREATE, os.ModePerm)
		}
	}
	return
}

func (s *RotateFileHandler) Close() error {
	return s.file.Close()
}
func checkLogger() {
	if handler != nil {
		_ = handler.Close()
		slog.SetDefault(slog.Default())
	}
	logFile := conf.GetString("log.file", "")
	c := NewConfig(conf)
	opt := new(slog.HandlerOptions)
	{
		opt.AddSource = conf.GetBoolean("log.source", true)
		lever := new(slog.Level)
		if err := lever.UnmarshalText([]byte(conf.GetString("log.level", "info"))); err == nil {
			opt.Level = lever
		} else {
			opt.Level = slog.LevelInfo
		}
	}
	if logFile == "" {
		log := slog.New(slog.NewJSONHandler(os.Stdout, opt))
		slog.SetDefault(log)
	} else {
		fn.Panic(os.MkdirAll(filepath.Dir(logFile), os.ModePerm))
		f := fn.Panic1(os.OpenFile(logFile, os.O_CREATE|os.O_APPEND, os.ModePerm))
		handler = &RotateFileHandler{
			path:    logFile,
			pattern: conf.GetString("log.pattern", "060102150405"),
			file:    f,
			limit:   c.GetByteSizeOr("log.size", big.NewInt(1024*1024*10)).Int64(),
			lock:    sync.Mutex{},
		}
		log := slog.New(slog.NewJSONHandler(handler, opt))
		slog.SetDefault(log)
	}
}

type adaptor struct {
}

func (a adaptor) Info(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.Info(f, v[1:]...)
			return
		}
	}
	slog.Info("", v...)
}

func (a adaptor) Infoln(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.Info(f, v[1:]...)
			return
		}
	}
	slog.Info("", v...)
}

func (a adaptor) Infof(format string, v ...any) {
	slog.Info(fmt.Sprintf(format, v...))
}

func (a adaptor) Warn(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.Warn(f, v[1:]...)
			return
		}
	}
	slog.Warn("", v...)
}

func (a adaptor) Warnln(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.Warn(f, v[1:]...)
			return
		}
	}
	slog.Warn("", v...)
}

func (a adaptor) Warnf(format string, v ...any) {
	slog.Warn(fmt.Sprintf(format, v))
}

func (a adaptor) Error(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.Error(f, v[1:]...)
			return
		}
	}
	slog.Error("", v...)
}

func (a adaptor) Errorln(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.Error(f, v[1:]...)
			return
		}
	}
	slog.Error("", v...)
}

func (a adaptor) Errorf(format string, v ...any) {
	slog.Error(format, v...)
}

func (a adaptor) InfoContext(ctx context.Context, v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.InfoContext(ctx, f, v[1:]...)
			return
		}
	}
	slog.InfoContext(ctx, "", v...)
}

func (a adaptor) InfoContextf(ctx context.Context, format string, v ...any) {
	slog.InfoContext(ctx, fmt.Sprintf(format, v...))
}

func (a adaptor) WarnContext(ctx context.Context, v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.WarnContext(ctx, f, v[1:]...)
			return
		}
	}
	slog.WarnContext(ctx, "", v...)
}

func (a adaptor) WarnContextf(ctx context.Context, format string, v ...any) {
	slog.WarnContext(ctx, fmt.Sprintf(format, v...))
}

func (a adaptor) ErrorContext(ctx context.Context, v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			slog.ErrorContext(ctx, f, v[1:]...)
			return
		}
	}
	slog.ErrorContext(ctx, "", v...)
}

func (a adaptor) ErrorContextf(ctx context.Context, format string, v ...any) {
	slog.ErrorContext(ctx, fmt.Sprintf(format, v...))
}
