package schedule

import (
	"time"

	"github.com/goexl/gox"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleOnce)(nil)

type scheduleOnce struct {

	params *params
	add    *addParams

	delay      time.Duration
}

func newScheduleOnce(params *params, add *addParams) *scheduleOnce {
	return &scheduleOnce{
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
