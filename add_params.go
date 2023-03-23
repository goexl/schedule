package schedule

import (
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
	name     string
}

func newAddParams(scheduler *Scheduler, worker worker) (add *addParams) {
	add = new(addParams)
	add.typ = typeImmediately
	add.schedule = newScheduleOnce(scheduler.params, add)
	add.worker = worker

	return
}

func (ap *addParams) checkLimit(scheduler *Scheduler) (err error) {
	if nil != ap.limit {
		err = ap.limit.check(scheduler)
	}

	return
}
