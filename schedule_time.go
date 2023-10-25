package schedule

import (
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
	"github.com/goexl/log"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleTime)(nil)

type scheduleTime struct {
	params *addParams
	self   *timeParams
	logger log.Logger
}

func newScheduleTime(params *addParams, self *timeParams, logger log.Logger) *scheduleTime {
	return &scheduleTime{
		params: params,
		self:   self,
		logger: logger,
	}
}

func (st *scheduleTime) Next(_ time.Time) (next time.Time) {
	next = rand.New().Time().Between(st.self.from, st.self.to).Build().Generate()
	fields := gox.Fields[any]{
		field.New("type", "random.time"),
		field.New("next", next),
	}
	if "" != st.params.name {
		fields.Add(field.New("name", st.params.name))
	}
	st.logger.Debug("任务调度", fields...)

	return
}
