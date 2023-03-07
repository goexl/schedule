package schedule

import (
	"time"

	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleOnce)(nil)

type scheduleOnce struct {
	executed bool
	id       *cron.EntryID
	cron     *cron.Cron
}

func newScheduleOnce(id *cron.EntryID, cron *cron.Cron) *scheduleOnce {
	return &scheduleOnce{
		executed: false,
		id:       id,
		cron:     cron,
	}
}

func (so *scheduleOnce) Next(from time.Time) (runtime time.Time) {
	// 只执行一次
	if so.executed {
		so.cron.Remove(*so.id)
	}

	if from.Before(time.Now()) {
		from = time.Now()
	}
	runtime = from.Add(100 * time.Millisecond)
	so.executed = true

	return
}
