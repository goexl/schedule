package schedule

type randomBuilder struct {
	params  *params
	add     *addParams
	builder *addBuilder
}

func newRandomBuilder(params *params, add *addParams, builder *addBuilder) *randomBuilder {
	return &randomBuilder{
		params:  params,
		add:     add,
		builder: builder,
	}
}

func (rb *randomBuilder) Duration() *durationBuilder {
	return newDurationBuilder(rb.params, rb.add, rb.builder)
}

func (rb *randomBuilder) Time() *timeBuilder {
	return newTimeBuilder(rb.params, rb.add, rb.builder)
}
