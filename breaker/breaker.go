package breaker

import (
	"errors"
	"sync"
	"time"
)

// Breaker simplified breaker. see original https://github.com/sony/gobreaker
type Breaker struct {
	state      State
	counter    Counter
	configure  Configure
	mutex      sync.Mutex
	generation uint64
	expiry     time.Time
}

func defaultStrip(counter *Counter) bool {
	return counter.ConsecutiveFailures > 5
}

// Configure current Breaker
func (s *Breaker) Configure(c func(configure *Configure)) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	c(&s.configure)
	if s.configure.Interval <= 0 {
		s.configure.Interval = time.Second
	}
	if s.configure.Timeout <= 0 {
		s.configure.Timeout = time.Minute
	}
	if s.configure.ReadyToTrip == nil {
		s.configure.ReadyToTrip = defaultStrip
	}
}

// Requests current requests
func (s *Breaker) Requests() uint32 {
	return s.counter.Requests
}
func (s *Breaker) State() State {
	return s.state
}
func (s *Breaker) Name() string {
	return s.configure.Name
}

func (s *Breaker) Prepare() (handle func(success bool), err error) {
	generation, err := s.pre()
	if err != nil {
		return nil, err
	}
	return func(success bool) {
		s.after(generation, success)
	}, nil
}

func (s *Breaker) currentState(now time.Time) (State, uint64) {
	switch s.state {
	case StateClosed:
		if !s.expiry.IsZero() && s.expiry.Before(now) {
			s.newGeneration(now)
		}
	case StateOpen:
		if s.expiry.Before(now) {
			s.setState(StateHalfOpen, now)
		}
	}
	return s.state, s.generation
}

func (s *Breaker) pre() (uint64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	state, generation := s.currentState(now)

	if state == StateOpen {
		return generation, ErrOpenState
	} else if state == StateHalfOpen && s.counter.Requests >= s.configure.MaxRequests {
		return generation, ErrTooManyRequests
	}
	s.counter.OnRequest()
	return generation, nil
}

func (s *Breaker) after(before uint64, success bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := time.Now()
	state, generation := s.currentState(now)
	if generation != before { // ignore
		return
	}
	if success {
		s.onSuccess(state, now)
	} else {
		s.onFailure(state, now)
	}
}
func (s *Breaker) newGeneration(now time.Time) {
	s.generation++
	s.counter.Reset()
	switch s.state {
	case StateClosed:
		if s.configure.Interval == 0 {
			s.expiry = zero
		} else {
			s.expiry = now.Add(s.configure.Interval)
		}
	case StateOpen:
		s.expiry = now.Add(s.configure.Timeout)
	default:
		s.expiry = zero
	}
}
func (s *Breaker) setState(state State, now time.Time) {
	if s.state == state {
		return
	}
	prev := s.state
	s.state = state
	s.newGeneration(now)
	if s.configure.OnStateChange != nil {
		s.configure.OnStateChange(s.configure.Name, prev, state)
	}
}
func (s *Breaker) onSuccess(state State, now time.Time) {
	switch state {
	case StateClosed:
		s.counter.OnSuccess()
	case StateHalfOpen:
		s.counter.OnSuccess()
		if s.counter.ConsecutiveSuccesses >= s.configure.MaxRequests {
			s.setState(StateClosed, now)
		}
	}
}
func (s *Breaker) onFailure(state State, now time.Time) {
	switch state {
	case StateClosed:
		s.counter.OnFailure()
		if s.configure.ReadyToTrip(&s.counter) {
			s.setState(StateOpen, now)
		}
	case StateHalfOpen:
		s.setState(StateOpen, now)
	}
}

type Configure struct {
	Name          string                                  //the breaker name
	MaxRequests   uint32                                  // half open state allowed requests
	Interval      time.Duration                           //counter reset interval on [StateClosed]
	Timeout       time.Duration                           //[StateOpen] timeout then switch to [StateHalfOpen]
	ReadyToTrip   func(counter *Counter) bool             //check of counter to trip, default is when reach five consecutive failures
	OnStateChange func(name string, from State, to State) //optional state monitor
}
type Counter struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

func (s *Counter) OnRequest() {
	s.Requests++
}
func (s *Counter) OnSuccess() {
	s.TotalSuccesses++
	s.ConsecutiveSuccesses++
	s.ConsecutiveFailures = 0
}
func (s *Counter) OnFailure() {
	s.TotalFailures++
	s.ConsecutiveSuccesses = 0
	s.ConsecutiveFailures++
}
func (s *Counter) Reset() {
	s.Requests = 0
	s.TotalSuccesses = 0
	s.TotalFailures = 0
	s.ConsecutiveSuccesses = 0
	s.ConsecutiveFailures = 0
}

//go:generate stringer -type=State
type State int32

const (
	StateOpen State = iota
	StateHalfOpen
	StateClosed
)

var (
	zero               time.Time
	ErrTooManyRequests = errors.New("too many requests")
	ErrOpenState       = errors.New("breaker is open")
)
