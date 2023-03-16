package schedule

import (
	"time"

	"github.com/robfig/cron/v3"
)

type addParams struct {
	typ      typ
	worker   worker
	ticker   ticker
	schedule cron.Schedule
	limit    *limitParams
	id       string
	unique   bool
	echo     bool
	delay    time.Duration
}

func newAddParams(worker worker) *addParams {
	return &addParams{
		typ:    typeImmediately,
		worker: worker,
	}
}

func (ap *addParams) checkLimit(scheduler *Scheduler) (err error) {
	if nil != ap.limit {
		err = ap.limit.check(scheduler)
	}

	return
}
