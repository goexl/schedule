package schedule

import (
	"time"
)

type timeParams struct {
	from time.Time
	to   time.Time
}

func newTimeParams() *timeParams {
	return &timeParams{
		to: time.Now().Add(time.Hour * 24 * 365),
	}
}
