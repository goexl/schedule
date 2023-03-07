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
		delay:  time.Second,
	}
}

func (ap *addParams) checkLimit(scheduler *Scheduler) (checked bool) {
	if nil == ap.limit {
		checked = true
	} else {
		checked = ap.limit.check(scheduler)
	}

	return
}
