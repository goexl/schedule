package schedule

import (
	"time"
)

type durationParams struct {
	from time.Duration
	to   time.Duration
}

func newDurationParams() *durationParams {
	return &durationParams{
		to: time.Hour,
	}
}
