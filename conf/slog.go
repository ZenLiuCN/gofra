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
	"runtime"
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
	i = adaptor{slog.Default()}
}

type adaptor struct {
	l *slog.Logger
}

func (a adaptor) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !a.l.Enabled(ctx, level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(4, pcs[:])
	record := slog.NewRecord(time.Now(), level, msg, pcs[0])
	record.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = a.l.Handler().Handle(ctx, record)
}
func (a adaptor) Info(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			a.log(nil, slog.LevelInfo, fmt.Sprintf(f, v[1:]...))
			return
		}
	}
	a.log(nil, slog.LevelInfo, "", v...)
}

func (a adaptor) Infof(format string, v ...any) {
	a.log(nil, slog.LevelInfo, fmt.Sprintf(format, v...))
}

func (a adaptor) Warn(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			a.log(nil, slog.LevelWarn, f, v[1:]...)
			return
		}
	}
	a.log(nil, slog.LevelWarn, "", v...)
}

func (a adaptor) Warnf(format string, v ...any) {
	a.log(nil, slog.LevelWarn, fmt.Sprintf(format, v))
}

func (a adaptor) Error(v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			a.log(nil, slog.LevelError, f, v[1:]...)
			return
		}
	}
	a.log(nil, slog.LevelError, "", v...)
}

func (a adaptor) Errorf(format string, v ...any) {
	a.log(nil, slog.LevelError, fmt.Sprintf(format, v...))
}

func (a adaptor) InfoContext(ctx context.Context, v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			a.log(ctx, slog.LevelInfo, f, v[1:]...)
			return
		}
	}
	a.log(ctx, slog.LevelInfo, "", v...)
}

func (a adaptor) InfoContextf(ctx context.Context, format string, v ...any) {
	a.log(ctx, slog.LevelInfo, fmt.Sprintf(format, v...))
}

func (a adaptor) WarnContext(ctx context.Context, v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			a.log(ctx, slog.LevelWarn, f, v[1:]...)
			return
		}
	}
	a.log(ctx, slog.LevelWarn, "", v...)
}

func (a adaptor) WarnContextf(ctx context.Context, format string, v ...any) {
	a.log(ctx, slog.LevelWarn, fmt.Sprintf(format, v...))
}

func (a adaptor) ErrorContext(ctx context.Context, v ...any) {
	if len(v) > 0 {
		if f, ok := v[0].(string); ok {
			a.log(ctx, slog.LevelError, f, v[1:]...)
			return
		}
	}
	a.log(ctx, slog.LevelError, "", v...)
}

func (a adaptor) ErrorContextf(ctx context.Context, format string, v ...any) {
	a.log(ctx, slog.LevelError, fmt.Sprintf(format, v...))
}
