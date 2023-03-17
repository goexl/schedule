package schedule

import (
	"time"
)

type durationBuilder struct {
	params  *params
	add     *addParams
	builder *addBuilder
	self    *durationParams
}

func newDurationBuilder(params *params, add *addParams, builder *addBuilder) *durationBuilder {
	return &durationBuilder{
		params:  params,
		builder: builder,
		add:     add,
		self:    newDurationParams(),
	}
}

func (db *durationBuilder) From(duration time.Duration) *durationBuilder {
	db.self.from = duration

	return db
}

func (db *durationBuilder) To(duration time.Duration) *durationBuilder {
	db.self.to = duration

	return db
}

func (db *durationBuilder) Between(from time.Duration, to time.Duration) *durationBuilder {
	db.self.from = from
	db.self.to = to

	return db
}

func (db *durationBuilder) Build() (builder *addBuilder) {
	db.builder.self.schedule = newScheduleDuration(db.add, db.self, db.params.logger)
	builder = db.builder

	return
}
