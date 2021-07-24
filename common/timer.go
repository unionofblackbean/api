package common

import "time"

type Timer struct {
	start time.Time
	end   time.Time
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	t.start = time.Now()
}

func (t *Timer) Stop() {
	t.end = time.Now()
}

func (t *Timer) Duration() time.Duration {
	return t.end.Sub(t.start)
}