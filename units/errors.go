package units

import (
	"encoding/json"
	"net/http"
	"time"
)

type InternalError string

func (i InternalError) Error() string {
	return string(i)
}
func (i InternalError) Is(target error) (ok bool) {
	_, ok = target.(InternalError)
	if !ok {
		_, ok = target.(*InternalError)
	}
	return
}

type ResponseError struct {
	Code    int
	Message string
}

func (i ResponseError) Error() string {
	return i.Message
}

type JsonError struct {
	Timestamp int64  `json:"timestamp"`
	Code      int    `json:"code"`
	Message   string `json:"message,omitempty"`
}

// JsonSafeHandleFunc recover and returns json error object.
//
// 400/404/401/500 will expose as http status
func JsonSafeHandleFunc(h http.HandlerFunc, logger func(format string, args ...any)) (s http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			switch e := recover().(type) {
			case error:
				if ex, ok := e.(ResponseError); ok {
					switch ex.Code {
					case http.StatusNotFound, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError:
						w.WriteHeader(ex.Code)
					}
					_ = json.NewEncoder(w).Encode(JsonError{
						Timestamp: time.Now().Unix(),
						Code:      ex.Code,
						Message:   ex.Message,
					})
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					_ = json.NewEncoder(w).Encode(JsonError{
						Timestamp: time.Now().Unix(),
						Code:      500,
						Message:   "",
					})
				}
				logger("handle %s", r.RequestURI, e)
			case nil:
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(JsonError{
					Timestamp: time.Now().Unix(),
					Code:      500,
					Message:   "",
				})
				logger("handle %s", r.RequestURI, e)
			}
		}()
		h(w, r)
	}
}

// TextSafeHandleFunc recover and returns http status code with text error message. Note all error code are treat as http status code
func TextSafeHandleFunc(h http.HandlerFunc, logger func(format string, args ...any)) (s http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			switch e := recover().(type) {
			case error:
				if ex, ok := e.(ResponseError); ok {
					w.WriteHeader(ex.Code)
					_, _ = w.Write([]byte(ex.Message))
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				logger("handle %s %#+v", r.RequestURI, e)
			case nil:
			default:
				w.WriteHeader(http.StatusInternalServerError)
				logger("handle %s %#+v", r.RequestURI, e)
			}
		}()
		h(w, r)
	}
}
