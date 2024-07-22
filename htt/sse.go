package htt

import (
	"bytes"
	"context"
	"github.com/ZenLiuCN/gofra/conf"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// SSE component to send Server-Send-Event to http client.
//
// A SSE can only use once. [SSE.Await] should only call inside [http.Handler]. it will block until all events are done.
// All send methods may return [io.EOF] if the SSE already shutdown.
type SSE interface {
	Close() error                             //close the SSE, returns [io.EOF] if already closed.
	Send(data string) error                   //alias of SendEventId.
	Ping() error                              //send a ping comment, returns [io.EOF] if already closed.
	SendEvent(event, data string) error       //alias of SendEventId.
	SendEventID(id, event, data string) error //send event to SSE client, empty string will not write to stream, returns [io.EOF] if already closed.
	Retry(mills int) error                    // send retry to SSE client, returns [io.EOF] if already closed.
	Raw() http.ResponseWriter                 // the internal response writer.
	Await(ctx context.Context) error          // Await with optional context inside [http.Handler] , returns [io.EOF] if SSE already closed or the client closed connection.
}
type sse struct {
	ch chan *msg
	http.ResponseWriter
}
type msg struct {
	ping  bool
	retry uint32
	event string
	id    string
	data  string
}

var (
	bLF    = []byte("\n")
	bData  = []byte("data: ")
	bId    = []byte("id: ")
	bEvent = []byte("event: ")
	bRetry = []byte("retry: ")
	bPing  = []byte(": ping\n\n")
)

func (s *sse) Close() error {
	if s.ch == nil {
		return io.EOF
	}
	close(s.ch)
	s.ch = nil
	return nil
}
func (s *sse) Raw() http.ResponseWriter {
	return s.ResponseWriter
}
func (s *sse) Ping() (err error) {
	return s.send(&msg{ping: true})
}
func (s *sse) SendEventID(id, event, data string) error {
	return s.send(&msg{id: id, event: event, data: data})
}
func (s *sse) Retry(mills int) error {
	return s.send(&msg{retry: uint32(mills)})
}

type flush interface {
	FlushError() error
}

func (s *sse) Await(ctx context.Context) (err error) {
	if s.ch == nil {
		return io.EOF
	}
	defer func() {
		close(s.ch)
		s.ch = nil
	}()
	if ctx != nil {
		return s.withContext(ctx)
	} else {
		return s.withoutContext()
	}
}
func (s *sse) withoutContext() (err error) {
	var buf = new(bytes.Buffer)
	var ok bool
	var m *msg
	for {
		select {
		case m, ok = <-s.ch:
			if !ok {
				conf.Internal().Warn("sse chan closed")
				return
			}
			err = s.write(m, buf)
			if err != nil {
				return
			}
			m = nil
		}
	}
}
func (s *sse) write(m *msg, buf *bytes.Buffer) (err error) {
	if m.ping {
		buf.Write(bPing)
	} else if m.retry > 0 {
		buf.Write(bRetry)
		buf.WriteString(strconv.Itoa(int(m.retry)))
		buf.Write(bLF)
	} else {
		if m.id != "" {
			buf.Write(bId)
			buf.WriteString(m.id)
			buf.Write(bLF)
		}
		if m.event != "" {
			buf.Write(bEvent)
			buf.WriteString(m.event)
			buf.Write(bLF)
		}
		if m.data != "" {
			data := m.data
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
	}
	_, err = buf.WriteTo(s)
	if err != nil {
		return
	}
	err = s.ResponseWriter.(flush).FlushError()
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			return io.EOF
		}
		return
	}
	buf.Reset()
	return
}
func (s *sse) withContext(ctx context.Context) (err error) {
	var buf = new(bytes.Buffer)
	var ok bool
	var m *msg
	for {
		select {
		case <-ctx.Done():
			conf.Internal().Warn("sse context closed")
			return
		case m, ok = <-s.ch:
			if !ok {
				conf.Internal().Warn("sse chan closed")
				return
			}
			err = s.write(m, buf)
			if err != nil {
				return
			}
			m = nil
		}
	}
}

func (s *sse) send(msg *msg) (err error) {
	if s.ch == nil {
		return io.EOF
	}
	log.Printf("Send %+v", msg)
	s.ch <- msg
	return
}
func (s *sse) Send(data string) (err error) {
	return s.SendEventID("", "", data)
}

func (s *sse) SendEvent(event, data string) (err error) {
	return s.SendEventID("", event, data)
}

// NewSSE create new SSE, should not send headers manually, this function will send Server-Send-Event headers.
func NewSSE(w http.ResponseWriter, bufferSize ...int) (s SSE) {
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("Transfer-Encoding", "chunked")
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("X-Accel-Buffering", "no")
	w.Header().Add("Cache-Control", "no-cache, must-revalidate")
	w.WriteHeader(200)
	var buf = 1
	if len(bufferSize) > 0 && bufferSize[0] > 0 {
		buf = bufferSize[0]
	}
	s = &sse{make(chan *msg, buf), w}
	_ = s.Ping()
	return
}
