package schedule

import (
	"time"
)

type fixedBuilder struct {
	scheduler *Scheduler
	params    *addParams
	self      *fixedParams
}

func newFixedBuilder(time time.Time, scheduler *Scheduler, params *addParams) *fixedBuilder {
	return &fixedBuilder{
		scheduler: scheduler,
		params:    params,
		self:      newFixedParams(time),
	}
}

func (fb *fixedBuilder) Random(duration time.Duration) *fixedBuilder {
	fb.self.random = true
	fb.self.max = duration

	return fb
}

func (fb *fixedBuilder) Min(duration time.Duration) *fixedBuilder {
	fb.self.random = true
	fb.self.min = duration

	return fb
}

func (fb *fixedBuilder) Between(min time.Duration, max time.Duration) *fixedBuilder {
	fb.self.random = true
	fb.self.min = min
	fb.self.max = max

	return fb
}

func (fb *fixedBuilder) Delay(duration time.Duration) *fixedBuilder {
	fb.self.max = duration

	return fb
}

func (fb *fixedBuilder) Build() (_add *add) {
	fb.params.ticker = newFixedTicker(fb.self)
	_add = newAdd(fb.scheduler, fb.params)

	return
}
