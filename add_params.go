package schedule

import (
	"time"
)

type addParams struct {
	typ    typ
	worker worker
	ticker ticker
	limit  *limitParams
	id     string
	unique bool
	delay  time.Duration
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
