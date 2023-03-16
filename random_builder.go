package schedule

type randomBuilder struct {
	builder *addBuilder
}

func newRandomBuilder(builder *addBuilder) *randomBuilder {
	return &randomBuilder{
		builder: builder,
	}
}

func (rb *randomBuilder) Duration() (builder *durationBuilder) {
	rb.builder.params.typ = typeRandomDuration
	builder = newDurationBuilder(rb.builder)

	return
}

func (rb *randomBuilder) Time() (builder *timeBuilder) {
	rb.builder.params.typ = typeRandomTime
	builder = newTimeBuilder(rb.builder)

	return
}
