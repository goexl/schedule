package schedule

import (
	"strings"
	"time"
)

var _ ticker = (*durationTicker)(nil)

type durationTicker struct {
	duration time.Duration
}

func newDurationTicker(duration time.Duration) *durationTicker {
	return &durationTicker{
		duration: duration,
	}
}

func (dt *durationTicker) tick() string {
	builder := new(strings.Builder)
	builder.WriteString(every)
	builder.WriteString(space)
	builder.WriteString(dt.duration.String())

	return builder.String()
}
