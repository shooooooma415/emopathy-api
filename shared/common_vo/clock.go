package common_vo

import (
	"testing"
	"time"
)

type Clock interface {
	Now() time.Time
	Since(t time.Time) time.Duration
}

type clock struct{}

var _ Clock = (*clock)(nil)

func SystemClock() *clock {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return time.Now().UTC()
}

func (c *clock) Since(t time.Time) time.Duration {
	return c.Now().Sub(t)
}

type StubClock struct {
	now *time.Time
}

var _ Clock = (*StubClock)(nil)

func NewStubClock(tb testing.TB) *StubClock {
	tb.Helper()
	return &StubClock{
		now: nil,
	}
}

func (s *StubClock) Now() time.Time {
	if s.now != nil {
		return *s.now
	}
	return time.Now().UTC()
}

func (s *StubClock) SetNow(t time.Time) {
	s.now = &t
}

func (c *StubClock) Since(t time.Time) time.Duration {
	return c.Now().Sub(t)
}
