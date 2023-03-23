package schedule

import (
	"time"

	"github.com/goexl/gox"
)

type addBuilder struct {
	scheduler *Scheduler
	self      *addParams
	params    *params
}

func newAddBuilder(scheduler *Scheduler, worker worker, params *params) *addBuilder {
	return &addBuilder{
		scheduler: scheduler,
		self:      newAddParams(scheduler, worker),
		params:    params,
	}
}

func (ab *addBuilder) Id(id any) *addBuilder {
	ab.self.id = gox.ToString(id)

	return ab
}

func (ab *addBuilder) Duration(duration time.Duration) *addBuilder {
	ab.self.typ = typeDuration
	ab.self.ticker = newDurationTicker(duration)

	return ab
}

func (ab *addBuilder) Cron(cron string) *addBuilder {
	ab.self.typ = typeCron
	ab.self.ticker = newCronTicker(cron)

	return ab
}

func (ab *addBuilder) Fixed(time time.Time) *addBuilder {
	ab.self.typ = typeFixed
	ab.self.schedule = newScheduleFixed(time)

	return ab
}

func (ab *addBuilder) Random() (builder *randomBuilder) {
	ab.self.typ = typeRandom
	builder = newRandomBuilder(ab.params, ab.self, ab)

	return
}

func (ab *addBuilder) Limit() *limitBuilder {
	return newLimitBuilder(ab)
}

func (ab *addBuilder) Unique() *addBuilder {
	ab.self.unique = true

	return ab
}

func (ab *addBuilder) Name(name string) *addBuilder {
	ab.self.name = name

	return ab
}

func (ab *addBuilder) Build() *add {
	return newAdd(ab.scheduler, ab.self)
}
