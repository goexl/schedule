package schedule

import (
	"time"

	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
	"github.com/goexl/simaqian"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleTime)(nil)

type scheduleTime struct {
	params *timeParams
	logger simaqian.Logger
}

func newScheduleTime(params *timeParams, logger simaqian.Logger) *scheduleTime {
	return &scheduleTime{
		params: params,
		logger: logger,
	}
}

func (sd *scheduleTime) Next(_ time.Time) (next time.Time) {
	next = rand.New().Time().Between(sd.params.from, sd.params.to).Build().Generate()
	sd.logger.Debug("任务高度", field.New("type", "time"), field.New("next", next))

	return
}
