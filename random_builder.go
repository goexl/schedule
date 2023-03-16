package schedule

type randomBuilder struct {
	builder *addBuilder
}

func newRandomBuilder(builder *addBuilder) *randomBuilder {
	return &randomBuilder{
		builder: builder,
	}
}

func (rb *randomBuilder) Duration() *durationBuilder {
	return newDurationBuilder(rb.builder)
}

func (rb *randomBuilder) Time() *timeBuilder {
	return newTimeBuilder(rb.builder)
}
