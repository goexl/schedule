package schedule

import (
	"time"
)

type timeBuilder struct {
	params  *params
	builder *addBuilder
	self    *timeParams
}

func newTimeBuilder(params *params, builder *addBuilder) *timeBuilder {
	return &timeBuilder{
		params:  params,
		builder: builder,
		self:    newTimeParams(),
	}
}

func (tb *timeBuilder) From(time time.Time) *timeBuilder {
	tb.self.from = time

	return tb
}

func (tb *timeBuilder) To(time time.Time) *timeBuilder {
	tb.self.to = time

	return tb
}

func (tb *timeBuilder) Between(from time.Time, to time.Time) *timeBuilder {
	tb.self.from = from
	tb.self.to = to

	return tb
}

func (tb *timeBuilder) Build() (builder *addBuilder) {
	tb.builder.self.schedule = newScheduleTime(tb.self, tb.params.logger)
	builder = tb.builder

	return
}
