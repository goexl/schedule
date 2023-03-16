package schedule

import (
	"time"
)

type durationBuilder struct {
	builder *addBuilder
	params  *durationParams
}

func newDurationBuilder(builder *addBuilder) *durationBuilder {
	return &durationBuilder{
		builder: builder,
		params:  newDurationParams(),
	}
}

func (db *durationBuilder) From(duration time.Duration) *durationBuilder {
	db.params.from = duration

	return db
}

func (db *durationBuilder) To(duration time.Duration) *durationBuilder {
	db.params.to = duration

	return db
}

func (db *durationBuilder) Build() *addBuilder {
	db.builder.params.schedule = newScheduleDuration()

	return db.builder
}
