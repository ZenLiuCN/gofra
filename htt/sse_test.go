package htt

import (
	"errors"
	"github.com/ZenLiuCN/fn"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestSSE(t *testing.T) {
	fn.Panic(http.ListenAndServe(":8082", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Header", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		sse := NewSSE(w)
		tick := time.NewTicker(time.Second)
		go func(tk *time.Ticker) {
			defer func() {
				tk.Stop()
			}()
		loop:
			for {
				select {
				case tm := <-tk.C:
					t.Logf("on tick %s", tm)
					err := sse.Send(tm.String())
					if err != nil {
						t.Logf("%s", err)
						break loop
					}
				}
			}
			t.Logf("quit ticker")
		}(tick)
		err := sse.Await(nil)
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}
		t.Logf("done")
	})))
}
func TestSSEContext(t *testing.T) {
	fn.Panic(http.ListenAndServe(":8082", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Header", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		sse := NewSSE(w)
		tick := time.NewTicker(time.Second)
		go func(tk *time.Ticker) {
			defer func() {
				tk.Stop()
			}()
		loop:
			for {
				select {
				case tm := <-tk.C:
					t.Logf("on tick %s", tm)
					err := sse.Send(tm.String())
					if err != nil {
						t.Logf("%s", err)
						break loop
					}
				}
			}
			t.Logf("quit ticker")
		}(tick)
		err := sse.Await(r.Context())
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}
		t.Logf("done")
	})))
}
