package schedule

import (
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
	"github.com/goexl/simaqian"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleDuration)(nil)

type scheduleDuration struct {
	params *addParams
	self   *durationParams
	logger simaqian.Logger
}

func newScheduleDuration(params *addParams, self *durationParams, logger simaqian.Logger) *scheduleDuration {
	return &scheduleDuration{
		params: params,
		self:   self,
		logger: logger,
	}
}

func (sd *scheduleDuration) Next(from time.Time) (next time.Time) {
	diff := rand.New().Duration().Between(sd.self.from, sd.self.to).Build().Generate()
	next = from.Add(diff)
	fields := gox.Fields[any]{
		field.New("type", "random.duration"),
		field.New("diff", diff.Truncate(time.Second)),
		field.New("next", next),
	}
	if "" != sd.params.name {
		fields.Add(field.New("name", sd.params.name))
	}
	sd.logger.Debug("任务调度", fields...)

	return
}
