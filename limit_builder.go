package schedule

type limitBuilder struct {
	params  *limitParams
	builder *addBuilder
}

func newLimitBuilder(builder *addBuilder) *limitBuilder {
	return &limitBuilder{
		params:  newLimitParams(),
		builder: builder,
	}
}

func (lb *limitBuilder) Cpu(percent float64) *limitBuilder {
	lb.params.cpu = percent

	return lb
}

func (lb *limitBuilder) Memory(percent float64) *limitBuilder {
	lb.params.memory = percent

	return lb
}

func (lb *limitBuilder) Process(count int) *limitBuilder {
	lb.params.process = count

	return lb
}

func (lb *limitBuilder) Max(max int) *limitBuilder {
	lb.params.max = max

	return lb
}

func (lb *limitBuilder) Build() (builder *addBuilder) {
	lb.builder.self.limit = lb.params
	builder = lb.builder

	return
}
