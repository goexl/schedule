package schedule

import (
	"time"

	"github.com/goexl/simaqian"
)

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Logger(logger simaqian.Logger) *builder {
	b.params.logger = logger

	return b
}

func (b *builder) Unique() *builder {
	b.params.unique = true

	return b
}

func (b *builder) Delay(delay time.Duration) *builder {
	b.params.delay = delay

	return b
}

func (b *builder) Build() *Scheduler {
	return newScheduler(b.params)
}
