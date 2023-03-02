package schedule

import (
	"time"

	"github.com/goexl/gox"
)

type addBuilder struct {
	scheduler *Scheduler
	params    *addParams
}

func newAddBuilder(scheduler *Scheduler, worker worker) *addBuilder {
	return &addBuilder{
		scheduler: scheduler,
		params:    newAddParams(worker),
	}
}

func (ab *addBuilder) Id(id any) *addBuilder {
	ab.params.id = gox.ToString(id)

	return ab
}

func (ab *addBuilder) Duration(duration time.Duration) *addBuilder {
	ab.params.ticker = newDurationTicker(duration)

	return ab
}

func (ab *addBuilder) Cron(cron string) *addBuilder {
	ab.params.ticker = newCronTicker(cron)

	return ab
}

func (ab *addBuilder) Fixed(time time.Time) *fixedBuilder {
	return newFixedBuilder(time, ab.scheduler, ab.params)
}

func (ab *addBuilder) Build() *add {
	return newAdd(ab.scheduler, ab.params)
}
