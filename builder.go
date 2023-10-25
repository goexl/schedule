package schedule

import (
	"github.com/goexl/log"
)

type builder struct {
	params *params
}

func newBuilder() *builder {
	return &builder{
		params: newParams(),
	}
}

func (b *builder) Logger(logger log.Logger) *builder {
	b.params.logger = logger

	return b
}

func (b *builder) Unique() *builder {
	b.params.unique = true

	return b
}

func (b *builder) Limit() *limitBuilder[*builder] {
	return newLimitBuilder(b)
}

func (b *builder) Build() *Scheduler {
	return newScheduler(b.params)
}
