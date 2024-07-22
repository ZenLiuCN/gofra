package htt

import (
	"github.com/ZenLiuCN/ote"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	"path/filepath"
)

// Spa [2]string{wwwRoot,indexFile}
type Spa [2]string

var (
	scope ote.TelemetryProviderFn = func() (scope string, opt []trace.SpanStartOption) {
		return "spa", nil
	}
)

func (h Spa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t, cx := ote.ByContext(r.Context(), scope); t != nil {
		cx, span := t.StartSpan("spa", cx, attribute.String("folder", h[0]), attribute.String("index", h[1]))
		defer func() {
			defer span.End()
			if r, ok := t.HandleRecover(recover()); ok {
				panic(r)
			}
		}()
		r = r.WithContext(cx)
	}
	path := filepath.Join(h[0], r.URL.Path)
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h[0], h[1])) // when not found or is dir,redirect to index
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h[0])).ServeHTTP(w, r)
}
