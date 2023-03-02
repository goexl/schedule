package schedule

import (
	"time"
)

type addParams struct {
	typ    typ
	worker worker
	ticker ticker

	cron         string
	time         time.Time
	delayMaxRand int64
	delay        time.Duration
	duration     time.Duration

	id string
}

func newAddParams(worker worker) *addParams {
	return &addParams{
		typ:          typeTime,
		worker:       worker,
		delayMaxRand: 100,
	}
}
