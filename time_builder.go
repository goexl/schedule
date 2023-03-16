package schedule

import (
	"time"
)

type timeBuilder struct {
	builder *addBuilder
	params  *timeParams
}

func newTimeBuilder(builder *addBuilder) *timeBuilder {
	return &timeBuilder{
		builder: builder,
		params:  newTimeParams(),
	}
}

func (tb *timeBuilder) From(time time.Time) *timeBuilder {
	tb.params.from = time

	return tb
}

func (tb *timeBuilder) To(time time.Time) *timeBuilder {
	tb.params.to = time

	return tb
}

func (tb *timeBuilder) Build() *addBuilder {
	tb.builder.params.schedule = newScheduleTime(tb.params)

	return tb.builder
}
