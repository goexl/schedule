package schedule

type limitBuilder struct {
	params  *limitParams
	builder *addBuilder
}

func newLimitBuilder() *limitBuilder {
	return &limitBuilder{
		params: newLimitParams(),
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
	lb.builder.params.limit = lb.params
	builder = lb.builder

	return
}
