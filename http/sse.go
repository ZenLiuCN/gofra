package http

import (
	"bytes"
	"github.com/ZenLiuCN/fn"
	"net/http"
	"strconv"
	"strings"
)

type SSE interface {
	Send(data string) error
	Ping() error
	SendEvent(event, data string) error
	SendEventID(id, event, data string) error
	Retry(mills int) error
	Raw() http.ResponseWriter
}
type sse [1]http.ResponseWriter

var (
	bLF    = []byte("\n")
	bData  = []byte("data: ")
	bId    = []byte("id: ")
	bEvent = []byte("event: ")
	bRetry = []byte("retry: ")
	bPing  = []byte(": ping\n\n")
)

func (s sse) Raw() http.ResponseWriter {
	return s[0]
}
func (s sse) Ping() (err error) {
	_, err = s[0].Write(bPing)
	return
}
func (s sse) SendEventID(id, event, data string) error {
	buf := fn.GetBuffer()
	defer func() {
		fn.PutBuffer(buf)
	}()
	if id != "" {
		buf.Write(bId)
		buf.WriteString(id)
		buf.Write(bLF)
	}
	if event != "" {
		buf.Write(bEvent)
		buf.WriteString(event)
		buf.Write(bLF)
	}
	if data != "" {
		i := strings.IndexByte(data, '\n')
		for i >= 0 {
			buf.Write(bData)
			buf.WriteString(data[:i])
			buf.Write(bLF)
			data = data[i+1:]
			i = strings.IndexByte(data, '\n')
		}
		if len(data) > 0 {
			buf.Write(bData)
			buf.WriteString(data)
			buf.Write(bLF)
		}
	}
	buf.Write(bLF)
	return s.send(buf)
}
func (s sse) Retry(mills int) error {
	buf := fn.GetBuffer()
	defer func() {
		fn.PutBuffer(buf)
	}()
	buf.Write(bRetry)
	buf.WriteString(strconv.Itoa(mills))
	buf.Write(bLF)
	return s.send(buf)
}

func (s sse) send(buf *bytes.Buffer) (err error) {
	_, err = buf.WriteTo(s[0])
	return
}
func (s sse) Send(data string) (err error) {

	return s.SendEventID("", "", data)
}

func (s sse) SendEvent(event, data string) (err error) {
	return s.SendEventID("", event, data)
}

// NewSSE create new SSE, should not send headers manually, this function will send Server-Send-Event headers.
func NewSSE(w http.ResponseWriter) SSE {
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("X-Accel-Buffering", "no")
	w.Header().Add("Cache-Control", "no-cache")
	w.WriteHeader(200)
	return &sse{w}
}
