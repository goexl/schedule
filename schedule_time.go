package schedule

import (
	"time"

	"github.com/goexl/gox/rand"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleTime)(nil)

type scheduleTime struct {
	params *timeParams
}

func newScheduleTime(params *timeParams) *scheduleTime {
	return &scheduleTime{
		params: params,
	}
}

func (sd *scheduleTime) Next(from time.Time) (next time.Time) {
	next = rand.New().Time().Between(sd.params.from, sd.params.to).Build().Generate()

	return
}
