package schedule

import (
	"time"
)

type fixedBuilder struct {
	builder *addBuilder
	params  *fixedParams
}

func newFixedBuilder(builder *addBuilder, time time.Time) *fixedBuilder {
	return &fixedBuilder{
		builder: builder,
		params:  newFixedParams(time),
	}
}

func (fb *fixedBuilder) Random(duration time.Duration) *fixedBuilder {
	fb.params.random = true
	fb.params.max = duration

	return fb
}

func (fb *fixedBuilder) Min(duration time.Duration) *fixedBuilder {
	fb.params.random = true
	fb.params.min = duration

	return fb
}

func (fb *fixedBuilder) Between(min time.Duration, max time.Duration) *fixedBuilder {
	fb.params.random = true
	fb.params.min = min
	fb.params.max = max

	return fb
}

func (fb *fixedBuilder) Delay(duration time.Duration) *fixedBuilder {
	fb.params.max = duration

	return fb
}

func (fb *fixedBuilder) Build() (builder *addBuilder) {
	fb.builder.params.ticker = newFixedTicker(fb.params)
	builder = fb.builder

	return
}
