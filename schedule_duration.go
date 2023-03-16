package schedule

import (
	"time"

	"github.com/goexl/gox/rand"
	"github.com/robfig/cron/v3"
)

var _ cron.Schedule = (*scheduleDuration)(nil)

type scheduleDuration struct {
	params *durationParams
}

func newScheduleDuration(params *durationParams) *scheduleDuration {
	return &scheduleDuration{
		params: params,
	}
}

func (sd *scheduleDuration) Next(from time.Time) (next time.Time) {
	diff := rand.New().Duration().Between(sd.params.from, sd.params.to).Build().Generate()
	next = from.Add(diff)

	return
}
