package schedule

import (
	"github.com/goexl/simaqian"
)

type builder struct {
	logger simaqian.Logger
}

func newBuilder() *builder {
	return &builder{
		logger: simaqian.Default(),
	}
}

func (b *builder) Logger(logger simaqian.Logger) *builder {
	b.logger = logger

	return b
}

func (b *builder) Build() *Scheduler {
	return newScheduler(b.logger)
}
