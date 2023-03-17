package schedule

import (
	"time"

	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
	"github.com/goexl/simaqian"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleDuration)(nil)

type scheduleDuration struct {
	params *durationParams
	logger simaqian.Logger
}

func newScheduleDuration(params *durationParams, logger simaqian.Logger) *scheduleDuration {
	return &scheduleDuration{
		params: params,
		logger: logger,
	}
}

func (sd *scheduleDuration) Next(from time.Time) (next time.Time) {
	diff := rand.New().Duration().Between(sd.params.from, sd.params.to).Build().Generate()
	next = from.Add(diff)
	sd.logger.Debug("任务高度", field.New("type", "duration"), field.New("diff", diff), field.New("next", next))

	return
}
