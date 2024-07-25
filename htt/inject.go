package htt

import (
	"bytes"
	"context"
	"github.com/ZenLiuCN/gofra/conf"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RouterConfigurer struct {
	*mux.Router
}

func RouterConfigurerOf(r *mux.Router) RouterConfigurer {
	return RouterConfigurer{r}
}
func (c RouterConfigurer) Get() *mux.Router {
	return c.Router
}

// WithTelemetry tracer provider is injected, not need in opts
func (c RouterConfigurer) WithTelemetry(service string, opts ...otelmux.Option) RouterConfigurer {
	opts = append(opts, otelmux.WithTracerProvider(otel.GetTracerProvider()))
	c.Use(otelmux.Middleware(service, opts...))
	return c
}

func parseCORS(cfg conf.Config) (h, o, a string) {
	if cfg != nil {
		headers := cfg.GetStringList("cors.headers")
		var b bytes.Buffer
		if len(headers) != 0 {
			for _, header := range headers {
				if b.Len() > 0 {
					b.WriteByte(',')
				}
				b.WriteString(header)
			}
		}
		if cfg.GetBoolean("cors.authorization", false) {
			if b.Len() > 0 {
				b.WriteByte(',')
			}
			b.WriteString("Authorization")
		}
		if b.Len() == 0 {
			b.WriteByte('*')
		}
		h = b.String()
		b.Reset()
		cc := cfg.GetBoolean("cors.credentials", false)
		if cc {
			a = "true"
		}
		origins := cfg.GetStringList("cors.origin")
		if len(origins) != 0 {
			for _, header := range origins {
				if b.Len() > 0 {
					b.WriteByte(',')
				}
				b.WriteString(header)
			}
		}
		if b.Len() == 0 {
			b.WriteByte('*')
		}
		o = b.String()
	} else {
		o = "*"
		h = "*"
	}
	return
}
func (c RouterConfigurer) WithCORS(cfg conf.Config) RouterConfigurer {
	if cfg == nil {
		return c
	}
	c.Use(mux.CORSMethodMiddleware(c.Router))
	h, o, a := parseCORS(cfg)
	c.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			/*			if req.Method == http.MethodOptions {
						if a!=""{
							w.Header().Set("Access-Control-Allow-Credentials", a)
						}
						w.Header().Set("Access-Control-Allow-Headers", h)
						w.Header().Set("Access-Control-Expose-Header", h)
						w.Header().Set("Access-Control-Allow-Origin", o)
					}*/
			if a != "" {
				w.Header().Set("Access-Control-Allow-Credentials", a)
			}
			w.Header().Set("Access-Control-Allow-Headers", h)
			w.Header().Set("Access-Control-Expose-Header", h)
			w.Header().Set("Access-Control-Allow-Origin", o)
			next.ServeHTTP(w, req)
		})
	})
	return c
}

// WithSPA at root path
// tpl: the routing prefix template
// folder: the local directory contains all SPA files
// index: the index file name
func (c RouterConfigurer) WithSPA(tpl, folder, index string) RouterConfigurer {
	c.PathPrefix(tpl).Handler(Spa([2]string{folder, index})).Name("pages")
	return c
}

// Launch see [StartServer]
func (c RouterConfigurer) Launch(name string, cfg conf.Config, configure func(server *http.Server), closerConsumer func(func())) {
	StartServer(name, c.Router, cfg, configure, closerConsumer)
}

// LaunchContextual see [StartContextServer]
func (c RouterConfigurer) LaunchContextual(name string, cfg conf.Config, timeout time.Duration, configure func(server *http.Server), closerConsumer func(func(context.Context))) {
	StartContextServer(name, c.Router, cfg, timeout, configure, closerConsumer)
}

/*
StartServer with mux.Router and [conf.Config].This function will block until [http.Server] is shutdown

HOCON sample:

	{
	address: "0.0.0.0:8080" # default address to listen with
	writeTimeout: 30s
	readTimeout: 30s
	idleTimeout: 30s
	keepAlive: false
	}
*/
func StartServer(name string, r *mux.Router, c conf.Config, configure func(server *http.Server), closerConsumer func(func())) {
	if c == nil {
		c = conf.Empty()
	}
	server := new(http.Server)
	server.Addr = conf.OrElse("address", "0.0.0.0:8080", c, c.GetString)
	server.Handler = r
	server.WriteTimeout = conf.OrElse("writeTimeout", time.Second*30, c, c.GetTimeDuration)
	server.ReadTimeout = conf.OrElse("readTimeout", time.Second*30, c, c.GetTimeDuration)
	server.IdleTimeout = conf.OrElse("idleTimeout", time.Second*60, c, c.GetTimeDuration)
	server.ErrorLog = log.Default()
	server.SetKeepAlivesEnabled(c.GetBoolean("keepAlive", false))
	if configure != nil {
		configure(server)
	}
	i := conf.Internal()
	shutdown := make(chan struct{})
	ch := make(chan os.Signal, 1)
	closer := func() {
		ch <- syscall.SIGINT
	}
	closerConsumer(closer)
	go func() {
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		<-ch
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			i.Errorf("http %s server shutting down %+v", name, err)
		} else {
			i.Infof("http %s server shutdown success", name)
			close(shutdown)
		}
		signal.Stop(ch)
	}()
	i.Infof("http %s server listen %s ", name, server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		i.Errorf("shutdown http %s server error %+v", name, err)
	}
	<-shutdown
	close(ch)
}

/*
StartContextServer with mux.Router and [conf.Config].This function will block
until [http.Server] is shutdown. This function also support a context for close
server.

HOCON sample:

	{
	address: "0.0.0.0:8080" # default address to listen with
	writeTimeout: 30s
	readTimeout: 30s
	idleTimeout: 30s
	keepAlive: false
	}
*/
func StartContextServer(name string, r *mux.Router, c conf.Config, timeout time.Duration, configure func(server *http.Server), closerConsumer func(func(ctx context.Context))) {
	if c == nil {
		c = conf.Empty()
	}
	server := new(http.Server)
	server.Addr = conf.OrElse("address", "0.0.0.0:8080", c, c.GetString)
	server.Handler = r
	server.WriteTimeout = conf.OrElse("writeTimeout", time.Second*30, c, c.GetTimeDuration)
	server.ReadTimeout = conf.OrElse("readTimeout", time.Second*30, c, c.GetTimeDuration)
	server.IdleTimeout = conf.OrElse("idleTimeout", time.Second*60, c, c.GetTimeDuration)
	server.ErrorLog = log.Default()
	server.SetKeepAlivesEnabled(c.GetBoolean("keepAlive", false))
	if configure != nil {
		configure(server)
	}
	i := conf.Internal()
	shutdown := make(chan struct{})
	ch := make(chan os.Signal, 1)
	ccx := make(chan context.Context, 1)
	closer := func(ctx context.Context) {
		ch <- syscall.SIGINT
		ccx <- ctx
	}
	closerConsumer(closer)
	go func() {
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		<-ch
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		var cx context.Context
	loop:
		for {
			select {
			case cx = <-ccx:
				break loop
			case <-ctx.Done():
				break loop
			}
		}
		if cx == nil {
			var cc context.CancelFunc
			ctx, cc = context.WithTimeout(context.Background(), timeout)
			defer cc()
		} else {
			ctx = cx
		}
		if err := server.Shutdown(ctx); err != nil {
			i.Errorf("http %s server shutting down %+v", name, err)
		} else {
			i.Infof("http %s server shutdown success", name)
			close(shutdown)
		}
		signal.Stop(ch)
	}()
	i.Infof("http %s server listen %s ", name, server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		i.Errorf("shutdown http %s server error %+v", name, err)
	}
	<-shutdown
	close(ch)
	close(ccx)
}
