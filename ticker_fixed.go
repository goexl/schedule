package schedule

import (
	"time"
)

var _ ticker = (*fixedTicker)(nil)

type fixedTicker struct {
	time time.Time
}

func newFixedTicker(time time.Time) *fixedTicker {
	return &fixedTicker{
		time: time,
	}
}

func (ft *fixedTicker) tick() string {
	return newRuntime(ft.time).spec()
}
