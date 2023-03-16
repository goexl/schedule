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

func (ab *addBuilder) Delay(delay time.Duration) *addBuilder {
	ab.params.delay = delay

	return ab
}

func (ab *addBuilder) Duration(duration time.Duration) *addBuilder {
	ab.params.typ = typeDuration
	ab.params.ticker = newDurationTicker(duration)

	return ab
}

func (ab *addBuilder) Cron(cron string) *addBuilder {
	ab.params.typ = typeCron
	ab.params.ticker = newCronTicker(cron)

	return ab
}

func (ab *addBuilder) Fixed(time time.Time) *fixedBuilder {
	return newFixedBuilder(ab, time)
}

func (ab *addBuilder) Random() *randomBuilder {
	return newRandomBuilder(ab)
}

func (ab *addBuilder) Limit() *limitBuilder {
	return newLimitBuilder(ab)
}

func (ab *addBuilder) Unique() *addBuilder {
	ab.params.unique = true

	return ab
}

func (ab *addBuilder) Build() *add {
	return newAdd(ab.scheduler, ab.params)
}
