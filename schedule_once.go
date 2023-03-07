package schedule

import (
	"fmt"
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
}

func newScheduleOnce(id *cron.EntryID, cron *cron.Cron, params *params, add *addParams) *scheduleOnce {
	return &scheduleOnce{
		id:         id,
		cron:       cron,
		params:     params,
		add:        add,
	}
}

func (so *scheduleOnce) Next(from time.Time) (next time.Time) {
	next = from.Add(gox.Ifx(0 != so.add.delay, func() time.Duration {
		return so.add.delay
	}, func() time.Duration {
		return gox.Ift(0 != so.params.delay, so.params.delay, 100*time.Millisecond)
	})).Add(so.delay)
	// 以50毫秒为步进值，逐步增加间隔，直到任务被执行
	so.delay += 50 * time.Millisecond

	return
}

func (so *scheduleOnce) completed() {
	// 删除原来的任务，确保不会再被执行
	so.cron.Remove(*so.id)
}
