package sse

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Emitter use to send SSE events
type Emitter interface {
	//Close manual close emitter
	Close() error
	//Raw fetch raw writer
	Raw() http.ResponseWriter
	//Use new writer, only when current emitter closed
	Use(w http.ResponseWriter) error
	//OnClose add on close hook, replace exists one
	OnClose(fn func())
	//Ping send a ping comment
	Ping() error
	//Retry send retry message
	Retry(mills int) error
	//Data send only data
	Data(data string) error
	//Event send event with id type and data, id and type can be empty.
	Event(id, evt, data string) error
	//Await block current goroutine until the request disconnected or Emitter been closed.
	//The context is required and should be relative to request context
	Await(ctx context.Context) error
}

// Notifier use to receive
type Notifier interface {
	Close() error                                //Close the notifier manually.
	Raw() *http.Response                         //Raw response
	Use(r *http.Response) error                  //Use the notifier again after close.
	OnClose(fn func())                           //OnClose hook, replace exists hook.
	OnPing(fn func()) error                      //OnPing handle ping, can only config before [Notifier.Await]
	OnRetry(fn func(mills int)) error            //OnRetry handle retry, can only config before [Notifier.Await]
	OnEvent(fn func(id, evt, data string)) error //OnEvent handle event, can only config before [Notifier.Await]
	Await(ctx context.Context) error             //Await process message until connection closed or context done
}

// Logger for print info
type Logger interface {
	Warn(message string)
	Warnf(message string, args ...any)
	Debug(message string)
	Debugf(message string, args ...any)
	Error(message string)
	Errorf(message string, args ...any)
}

var (
	// ErrClosed note the closeable already closed
	ErrClosed = errors.New("already closed")
	// ErrContextRequired when context parameter is missing
	ErrContextRequired = errors.New("context is required")
	// ErrConnFailure connection failure
	ErrConnFailure = errors.New("connection fail")
	// ErrRunning when the reuse entity is working
	ErrRunning = errors.New("is running")
)

type dto struct {
	retry int
	id    string
	event string
	data  string
}

func (d *dto) empty() bool {
	return d.retry == 0 && d.data == ""
}
func (d *dto) retrying() bool {
	return d.retry > 0
}
func (d *dto) ping() bool {
	return d.retry == -1 && d.data == ""
}
func (d *dto) full() bool {
	return d.data != "" && d.id != "" && d.event != ""
}
func (d *dto) simple() bool {
	return d.data != "" && d.id == "" && d.event == ""
}
func (d *dto) typed() bool {
	return d.data != "" && d.id == "" && d.event != ""
}

//region emitter

type emitter struct {
	buf     int
	ch      chan *dto
	onClose func()
	log     Logger
	http.ResponseWriter
}

func NewEmitter(optLog Logger, w http.ResponseWriter, buf ...int) Emitter {
	x := &emitter{ResponseWriter: w}
	if len(buf) > 0 && buf[0] > 1 {
		x.buf = buf[0]
	} else {
		x.buf = 1
	}
	x.log = optLog
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("Transfer-Encoding", "chunked")
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("X-Accel-Buffering", "no")
	w.Header().Add("Cache-Control", "no-cache, must-revalidate")
	w.WriteHeader(200)
	return x
}

func (s *emitter) Close() error {
	if s.ch != nil {
		close(s.ch)
		if s.onClose != nil {
			s.onClose()
			s.onClose = nil
		}
		//! should send chunked end message?
		return nil
	}
	return ErrClosed
}
func (s *emitter) Use(w http.ResponseWriter) error {
	if s.ch != nil {
		return ErrRunning
	}
	s.ResponseWriter = w
	return nil
}
func (s *emitter) Raw() http.ResponseWriter {
	return s.ResponseWriter
}

func (s *emitter) OnClose(fn func()) {
	s.onClose = fn
}

func (s *emitter) Ping() error {
	if s.ch == nil {
		return ErrClosed
	}
	s.ch <- &dto{retry: -1}
	return nil
}

func (s *emitter) Retry(mills int) error {
	if s.ch == nil {
		return ErrClosed
	}
	s.ch <- &dto{retry: mills}
	return nil
}

func (s *emitter) Data(data string) error {
	if s.ch == nil {
		return ErrClosed
	}
	s.ch <- &dto{data: data}
	return nil
}

func (s *emitter) Event(id, evt, data string) error {
	if s.ch == nil {
		return ErrClosed
	}
	s.ch <- &dto{id: id, event: evt, data: data}
	return nil
}

type flush interface {
	FlushError() error
	io.Writer
}

func (s *emitter) send(out flush, m string, args ...any) (err error) {
	_, err = fmt.Fprintf(out, m, args...)
	if err != nil {
		if s.log != nil {
			s.log.Errorf("send failure: %s ", err)
		}
		return errors.Join(ErrConnFailure, err)
	}
	err = out.FlushError()
	if err != nil {
		if s.log != nil {
			s.log.Errorf("flush failure: %s ", err)
		}
		return errors.Join(ErrConnFailure, err)
	}
	return nil
}
func (s *emitter) make() {
	if s.ch == nil {
		s.ch = make(chan *dto, s.buf)
	}
}
func (s *emitter) Await(ctx context.Context) (err error) {
	if ctx != nil {
		return ErrContextRequired
	}
	s.make()
	var d *dto
	var ok bool
	var w = s.ResponseWriter.(flush)
	for {
		select {
		case d, ok = <-s.ch:
			if !ok {
				if s.log != nil {
					s.log.Debugf("closing")
				}
				goto end
			}
			if s.log != nil {
				s.log.Debugf("process event: %#+v", d)
			}
			switch {
			case d.ping():
				err = s.send(w, ": ping\n\n")
				if err != nil {
					goto fail
				}
			case d.retrying():
				err = s.send(w, "retry:%d\n\n", d.retry)
				if err != nil {
					goto fail
				}
			case d.full():
				err = s.send(w, "id: %s\nevent: %s\ndata: %s\n\n", d.id, d.event, d.data)
				if err != nil {
					goto fail
				}
			case d.typed():
				err = s.send(w, "event: %s\ndata: %s\n\n", d.event, d.data)
				if err != nil {
					goto fail
				}
			case d.simple():
				err = s.send(w, "data: %s\n\n", d.event, d.data)
				if err != nil {
					goto fail
				}
			default:
				if s.log != nil {
					s.log.Warnf("send nothing for %#+v", d)
				}
			}
			d = nil
		case <-ctx.Done():
			if s.log != nil {
				s.log.Debugf("request context close")
			}
			close(s.ch)
			goto end
		}
	}
fail:
	close(s.ch)
	s.ch = nil
	if s.onClose != nil {
		s.onClose()
		s.onClose = nil
	}
	return err
end:
	s.ch = nil
	return ErrClosed
}

//endregion emitter

//region notifier

type notifier struct {
	dataBuf uint32
	buf     int
	cc      func()
	onClose func()
	onRetry func(mills int)
	onPing  func()
	onEvent func(id, evt, data string)
	log     Logger
	*http.Response
}

func clamp(a, b, v int) int {
	if a > b {
		c := b
		b = a
		a = c
	}
	if v >= b {
		return b
	}
	if v <= a {
		return a
	}
	return v
}

// NewNotifier make new notifier. optLog is optional [Logger]; buf is data
// channel buffer between 1~5; dataBuf is message buffer between 512B~8KB
func NewNotifier(optLog Logger, res *http.Response, buf int, dataBuf int) Notifier {
	x := &notifier{Response: res, log: optLog}
	x.buf = clamp(1, 5, buf)
	x.dataBuf = uint32(clamp(512, 1024*8, dataBuf))
	return x
}
func FillSSERequestHeader(w http.Header) {
	w.Set("Connection", "keep-alive")
	w.Set("Accept", "text/event-stream")
	w.Set("Cache-Control", "no-cache")
}
func (s *notifier) Close() error {
	if s.cc == nil {
		return ErrClosed
	}
	s.cc()
	s.cc = nil
	if s.onClose != nil {
		s.onClose()
		s.onClose = nil
	}
	return nil
}

func (s *notifier) Raw() *http.Response {
	return s.Response
}

func (s *notifier) Use(r *http.Response) error {
	if s.cc != nil {
		return ErrRunning
	}
	s.Response = r
	return nil
}

func (s *notifier) OnClose(fn func()) {
	s.onClose = fn
}

func (s *notifier) OnPing(fn func()) error {
	if s.cc != nil {
		return ErrRunning
	}
	s.onPing = fn
	return nil
}

func (s *notifier) OnRetry(fn func(mills int)) error {
	if s.cc != nil {
		return ErrRunning
	}
	s.onRetry = fn
	return nil
}

func (s *notifier) OnEvent(fn func(id string, evt string, data string)) error {
	if s.cc != nil {
		return ErrRunning
	}
	s.onEvent = fn
	return nil
}

var (
	ping        = []byte(":ping")
	headerID    = []byte("id:")
	headerData  = []byte("data:")
	headerEvent = []byte("event:")
	headerRetry = []byte("retry:")
)

func (s *notifier) Await(ctx context.Context) error {
	if s.cc != nil {
		return ErrRunning
	}
	if ctx == nil {
		return ErrContextRequired
	}
	ctx, s.cc = context.WithCancel(ctx)
	bins, errs := s.goRead(ctx)
	defer func() {
		close(bins)
		close(errs)
		if s.cc != nil {
			s.cc()
			s.cc = nil
		}
	}()
	trimHeader := func(size int, data []byte) []byte {
		if data == nil || len(data) < size {
			return data
		}
		data = data[size:]
		if len(data) > 0 && data[0] == 32 {
			data = data[1:]
		}
		if len(data) > 0 && data[len(data)-1] == 10 {
			data = data[:len(data)-1]
		}
		return data
	}
	var d = &dto{}
	var bin []byte
	var err error
	var data []byte
	for {
		select {
		case bin = <-bins:
			if s.cc == nil {
				goto closed
			}
			n := len(bin)
			if n == 1 {
				continue
			}
			for _, line := range bytes.FieldsFunc(bin, func(r rune) bool { return r == '\n' || r == '\r' }) {
				switch {
				case bytes.Equal(line, ping):
					d.retry = -1
					goto proc
				case bytes.HasPrefix(line, headerID):
					d.id = string(trimHeader(len(headerID), line))
				case bytes.HasPrefix(line, headerData):
					data = append(data, append(trimHeader(len(headerData), line), byte('\n'))...)
				case bytes.Equal(line, bytes.TrimSuffix(headerData, []byte(":"))):
					data = append(data, byte('\n'))
				case bytes.HasPrefix(line, headerEvent):
					d.event = string(trimHeader(len(headerEvent), line))
				case bytes.HasPrefix(line, headerRetry):
					res := string(trimHeader(len(headerRetry), line))
					d.retry, err = strconv.Atoi(res)
					if err != nil {
						if s.log != nil {
							s.log.Warnf("bad retry data: %s", res)
						}
						d.retry = 0
						err = nil
					}
				}
			}
			d.data = string(data)
			if cap(data) < 512 { //! max 512 buffer
				data = data[:0]
			} else {
				data = nil
			}
		proc:
			switch {
			case d.ping():
				if s.onPing != nil {
					s.onPing()
				}
				if s.log != nil {
					s.log.Debugf("on ping")
				}
			case d.retrying():
				if s.onRetry != nil {
					s.onRetry(d.retry)
				}
				if s.log != nil {
					s.log.Debugf("on retry: %d", d.retry)
				}
			case d.empty():
				if s.log != nil {
					s.log.Warnf("empty message received")
				}
			default:
				if s.onEvent != nil {
					s.onEvent(d.id, d.event, d.data)
				}
				if s.log != nil {
					s.log.Debugf("message received:#+v", d)
				}
			}
		case err = <-errs:
			if err == io.EOF {
				if s.log != nil {
					s.log.Debugf("read empty chunk")
				}
				continue
			} else if err == context.DeadlineExceeded {
				if s.log != nil {
					s.log.Debugf("context closed")
				}
				goto end
			}
			if s.log != nil {
				s.log.Errorf("connection error %s", err)
			}
			goto err
		}
	}
closed:
	if s.onClose != nil {
		c := s.onClose
		s.onClose = nil
		c()
	}
	return ErrClosed
end:
	if s.onClose != nil {
		c := s.onClose
		s.onClose = nil
		c()
	}
	return nil
err:
	if s.onClose != nil {
		c := s.onClose
		s.onClose = nil
		c()
	}
	return err
}
func chunkSplit(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	minial := func(a, b int) int {
		if a < 0 {
			return b
		}
		if b < 0 {
			return a
		}
		if a > b {
			return b
		}
		return a
	}
	chunkEnded := func(data []byte) (int, int) {
		irr := bytes.Index(data, []byte("\r\r"))
		inn := bytes.Index(data, []byte("\n\n"))
		irnn := bytes.Index(data, []byte("\r\n\n"))
		inrn := bytes.Index(data, []byte("\n\r\n"))
		irnrn := bytes.Index(data, []byte("\r\n\r\n"))
		st := minial(irr, minial(inn, minial(irnn, minial(inrn, irnrn))))
		n := 2
		if st == irnrn {
			n = 4
		} else if st == irnn || st == inrn {
			n = 3
		}
		return st, n
	}
	if i, n := chunkEnded(data); i >= 0 {
		return i + n, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
func (s *notifier) goRead(ctx context.Context) (out chan []byte, err chan error) {
	defer s.Response.Body.Close()
	var sc = bufio.NewScanner(s.Response.Body)
	sc.Buffer(make([]byte, 1024), int(s.dataBuf))
	sc.Split(chunkSplit)
	out = make(chan []byte, s.buf)
	err = make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				if s.log != nil {
					s.log.Debugf("context done")
				}
				err <- context.DeadlineExceeded
				return
			default:
				if sc.Scan() {
					out <- sc.Bytes()
				}
				if er := sc.Err(); err != nil {
					if er == context.Canceled {
						if s.log != nil {
							s.log.Debugf("connect fail")
						}
						err <- io.EOF
					} else {
						err <- er
					}
					return
				}
			}
		}
	}()
	return
}

//endregion notifier
