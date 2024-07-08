//go:build slog

package conf

import (
	"github.com/ZenLiuCN/fn"
	"log/slog"
	"math/big"
	"os"
	"path/filepath"
	"sync"
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
