package conf

import (
	"encoding/hex"
	"fmt"
	"github.com/ZenLiuCN/fn"
	hocon "github.com/go-akka/configuration"
	ho "github.com/go-akka/configuration/hocon"
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
	file    string
	conf    *hocon.Config
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

// Initialize with config file and [slog].
//
// HOCON sample for log config.
//
//	log{
//	 file: "path_of_log_file_with__name"
//	 pattern: "rotation_pattern_for_time"
//	 size: 2m
//	}
func Initialize(confFile string) {
	file = confFile
	conf = hocon.LoadConfig(confFile)
	checkLogger()
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

type (
	Config interface {
		fmt.Stringer
		GetObject(path string) Config
		GetObjects(path string) (r []Config)
		GetStringMap(path string) map[string]Config

		GetBoolean(path string, defaultVal ...bool) bool
		GetByteSize(path string) *big.Int
		GetByteSizeOr(path string, defaultVal ...*big.Int) *big.Int
		GetInt32(path string, defaultVal ...int32) int32
		GetInt64(path string, defaultVal ...int64) int64
		GetString(path string, defaultVal ...string) string
		GetFloat32(path string, defaultVal ...float32) float32
		GetFloat64(path string, defaultVal ...float64) float64
		GetTimeDuration(path string, defaultVal ...time.Duration) time.Duration
		GetTimeDurationInfiniteNotAllowed(path string, defaultVal ...time.Duration) time.Duration
		GetBooleanList(path string) []bool
		GetFloat32List(path string) []float32
		GetFloat64List(path string) []float64
		GetInt32List(path string) []int32
		GetInt64List(path string) []int64
		GetByteList(path string) []byte
		GetStringList(path string) []string
		HasPath(path string) bool
		IsObject(path string) bool
		IsArray(path string) bool

		RequiredString(path string) string
		RequiredInt32(path string) int32
		RequiredBoolean(path string) bool
		RequiredInt64(path string) int64

		ExistsString(path string, act func(string))
		ExistsBoolean(path string, act func(bool))
		ExistsInt32(path string, act func(int32))
		ExistsDuration(path string, act func(duration time.Duration))
		ExistsFloat64(path string, act func(d float64))

		GetTextMap(path string) map[string]string
	}
	config struct {
		*hocon.Config
	}
)

func (c config) GetTextMap(path string) (m map[string]string) {
	if c.HasPath(path) {
		v := c.GetValue(path)
		if v.IsObject() {
			m = make(map[string]string, len(v.GetObject().Items()))
			for s, value := range v.GetObject().Items() {
				m[s] = value.GetString()
			}
		}
	}
	return
}
func (c config) RequiredInt64(path string) int64 {
	return Required(path, c, c.GetInt64)
}
func (c config) RequiredBoolean(path string) bool {
	return Required(path, c, c.GetBoolean)
}
func (c config) RequiredInt32(path string) int32 {
	return Required(path, c, c.GetInt32)
}
func (c config) RequiredString(path string) string {
	return Required(path, c, c.GetString)
}
func (c config) ExistsFloat64(path string, act func(d float64)) {
	Exists(path, c, c.GetFloat64, act)
}
func (c config) ExistsDuration(path string, act func(duration time.Duration)) {
	Exists(path, c, c.GetTimeDurationInfiniteNotAllowed, act)
}
func (c config) ExistsString(path string, act func(string)) {
	Exists(path, c, c.GetString, act)
}
func (c config) ExistsBoolean(path string, act func(bool)) {
	Exists(path, c, c.GetBoolean, act)
}
func (c config) ExistsInt32(path string, act func(int32)) {
	Exists(path, c, c.GetInt32, act)
}
func (c config) GetStringMap(path string) map[string]Config {
	if c.HasPath(path) {
		n := c.GetNode(path)
		if n.IsObject() {
			o := n.GetObject()
			m := make(map[string]Config, len(o.GetKeys()))
			for _, s := range o.GetKeys() {
				m[s] = NewConfigOfValue(o.GetKey(s))
			}
			return m
		} else {
			return nil
		}
	} else {
		return nil
	}
}
func (c config) GetObject(path string) Config {
	if c.HasPath(path) {
		return config{
			c.GetConfig(path),
		}
	} else {
		return nil
	}
}
func (c config) GetByteSizeOr(path string, defaultVal ...*big.Int) *big.Int {
	if c.GetNode(path) != nil {
		return c.GetByteSize(path)
	} else if len(defaultVal) > 0 {
		return defaultVal[0]
	} else {
		return nil
	}
}
func (c config) GetObjects(path string) (r []Config) {
	if !c.HasPath(path) {
		return
	}
	a := c.GetNode(path).GetArray()
	for _, value := range a {
		if value.IsObject() {
			r = append(r, config{hocon.NewConfigFromRoot(ho.NewHoconRoot(value))})
		}
	}
	return
}

func NewConfig(c *hocon.Config) Config {
	return &config{c}
}
func NewConfigOfValue(c *ho.HoconValue) Config {
	return &config{Config: hocon.NewConfigFromRoot(ho.NewHoconRoot(c))}
}

func Exists[T any](path string, c Config, get func(path string, def ...T) T, consume func(T)) {
	if v, ok := c.(config); ok {
		if v.GetNode(path) != nil {
			consume(get(path))
		}
	} else {
		panic("invalid config instance")
	}

}

func ExistsMapping[T, V any](path string, c Config, get func(path string, def ...T) T, mapper func(T) V, consume func(V)) {
	if v, ok := c.(config); ok {
		if v.GetNode(path) != nil {
			consume(mapper(get(path)))
		}
	} else {
		panic("invalid config instance")
	}
}
func Required[T any](path string, c Config, get func(path string, def ...T) T) T {
	if v, ok := c.(config); ok {
		if v.GetNode(path) != nil {
			return get(path)
		} else {
			panic("missing configurer value of " + path)
		}
	} else {
		panic("invalid config instance")
	}
}
func OrElse[T any](path string, def T, c Config, get func(path string, def ...T) T) T {
	if v, ok := c.(config); ok {
		if v.GetNode(path) != nil {
			return get(path)
		} else {
			return def
		}
	} else {
		panic("invalid config instance")
	}
}
func GetConfig() Config {
	return config{conf}
}
func ReloadConfigurer(otherFile string) Config {
	defer checkLogger()
	if otherFile != "" {
		file = otherFile
	}
	conf = hocon.LoadConfig(file)
	return GetConfig()
}
func FlushConfigurer(data []byte) (c Config, success bool) {
	backupFile := fmt.Sprintf("%s.%d", file, time.Now().UnixMilli())
	defer func() {
		if r := recover(); r != nil {
			slog.With(slog.String("config", hex.EncodeToString(data))).Error("flush config fail", "error", r)
			fn.Panic(os.Rename(backupFile, file))
			success = false
			c = ReloadConfigurer("")
		}
		checkLogger()
	}()
	fn.Panic(os.Rename(file, backupFile))
	fn.Panic(os.WriteFile(file, data, os.ModePerm))
	return ReloadConfigurer(""), true
}
