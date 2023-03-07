package schedule

import (
	"math"
	"time"

	"github.com/goexl/gox"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleOnce)(nil)

type scheduleOnce struct {
	id     *cron.EntryID
	cron   *cron.Cron
	params *params
	add    *addParams

	delay      time.Duration
	_completed bool
}

func newScheduleOnce(id *cron.EntryID, cron *cron.Cron, params *params, add *addParams) *scheduleOnce {
	return &scheduleOnce{
		_completed: false,
		id:         id,
		cron:       cron,
		params:     params,
		add:        add,
	}
}

func (so *scheduleOnce) Next(from time.Time) (next time.Time) {
	// 只执行一次
	if so._completed {
		// 调度到一个完全不可能得到执行的时间
		next = time.Now().Add(math.MaxInt64)
		// 删除原来的任务，确保不会再被执行
		go so.cron.Remove(*so.id)
	} else {
		next = so.next(from)
	}

	return
}

func (so *scheduleOnce) completed() {
	so._completed = true
}

func (so *scheduleOnce) next(from time.Time) (next time.Time) {
	if from.Before(time.Now()) {
		from = time.Now()
	}
	next = from.Add(gox.Ifx(0 != so.add.delay, func() time.Duration {
		return so.add.delay
	}, func() time.Duration {
		return gox.Ift(0 != so.params.delay, so.params.delay, 100*time.Millisecond)
	})).Add(so.delay)
	so.delay += 100 * time.Millisecond

	return
}
