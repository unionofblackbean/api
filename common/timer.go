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
	t.start = NowUTC()
}

func (t *Timer) Stop() {
	t.end = NowUTC()
}

func (t *Timer) Duration() time.Duration {
	return t.end.Sub(t.start)
}
