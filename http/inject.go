package http

import (
	"bytes"
	"context"
	"github.com/ZenLiuCN/goinfra/conf"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Configurer [1]*mux.Router

func ConfigurerOf(r *mux.Router) Configurer {
	return Configurer([1]*mux.Router{r})
}
func (c Configurer) Get() *mux.Router {
	return c[0]
}
func (c Configurer) WithTelemetry(operation string, opts ...otelhttp.Option) Configurer {
	c[0].Use(otelhttp.NewMiddleware(operation, opts...))
	return c
}
func (c Configurer) WithCORS(cfg conf.Config) Configurer {
	c[0].Use(mux.CORSMethodMiddleware(c[0]))
	var h, o string
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
	c[0].Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Headers", h)
				w.Header().Set("Access-Control-Allow-Origin", o)
			}
			next.ServeHTTP(w, req)
		})
	})
	return c
}

// WithSPA at root path
// folder: the local directory contains all SPA files
// index: the index file name
func (c Configurer) WithSPA(folder, index string) Configurer {
	c[0].PathPrefix("/").Handler(Spa([2]string{folder, index}))
	return c
}

// Launch see [StartServer]
func (c Configurer) Launch(cfg conf.Config, closerConsumer func(func())) {
	StartServer(c[0], cfg, closerConsumer)
}

/*
StartServer with mux.Router and [conf.Config].This method will block until [http.Server] is shutdown

HOCON sample:

	{
	address: "0.0.0.0:8080" # default address to listen with
	writeTimeout: 30s
	readTimeout: 30s
	idleTimeout: 30s
	keepAlive: false
	}
*/
func StartServer(r *mux.Router, c conf.Config, closerConsumer func(func())) {
	server := new(http.Server)
	server.Addr = conf.OrElse("address", "0.0.0.0:8080", c, c.GetString)
	server.Handler = r
	server.WriteTimeout = conf.OrElse("writeTimeout", time.Second*30, c, c.GetTimeDuration)
	server.ReadTimeout = conf.OrElse("readTimeout", time.Second*30, c, c.GetTimeDuration)
	server.IdleTimeout = conf.OrElse("idleTimeout", time.Second*60, c, c.GetTimeDuration)
	server.ErrorLog = log.Default()
	server.SetKeepAlivesEnabled(c.GetBoolean("keepAlive", false))
	shutdown := make(chan struct{})
	ch := make(chan os.Signal, 1)
	closer := func() {
		ch <- syscall.SIGINT
	}
	closerConsumer(closer)
	go func() {

		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		<-ch
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("http server shutting down", err)
		} else {
			slog.Info("http server shutdown successfully")
			close(shutdown)
		}
	}()
	slog.Info("http server listen", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("shutdown http server", err)
	}
	<-shutdown
}
