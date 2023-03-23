package schedule

import (
	"time"

	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleFixed)(nil)

type scheduleFixed struct {
	time  time.Time
	delay time.Duration
}

func newScheduleFixed(time time.Time) *scheduleFixed {
	return &scheduleFixed{
		time: time,
	}
}

func (sf *scheduleFixed) Next(from time.Time) (next time.Time) {
	if from.After(sf.time) {
		sf.time = from
	}
	next = sf.time.Add(sf.delay)
	// 以50毫秒为步进值，逐步增加间隔，直到任务被执行
	sf.delay += 50 * time.Millisecond

	return
}
