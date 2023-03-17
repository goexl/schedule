package schedule

type randomBuilder struct {
	params  *params
	builder *addBuilder
}

func newRandomBuilder(params *params, builder *addBuilder) *randomBuilder {
	return &randomBuilder{
		params:  params,
		builder: builder,
	}
}

func (rb *randomBuilder) Duration() *durationBuilder {
	return newDurationBuilder(rb.params, rb.builder)
}

func (rb *randomBuilder) Time() *timeBuilder {
	return newTimeBuilder(rb.params, rb.builder)
}
