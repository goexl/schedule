package schedule

type limitBuilder[T limitType] struct {
	params  *limitParams
	builder T
}

func newLimitBuilder[T limitType](builder T) *limitBuilder[T] {
	return &limitBuilder[T]{
		params:  newLimitParams(),
		builder: builder,
	}
}

func (lb *limitBuilder[T]) Cpu(percent float64) *limitBuilder[T] {
	lb.params.cpu = percent

	return lb
}

func (lb *limitBuilder[T]) Memory(percent float64) *limitBuilder[T] {
	lb.params.memory = percent

	return lb
}

func (lb *limitBuilder[T]) Process(count int) *limitBuilder[T] {
	lb.params.process = count

	return lb
}

func (lb *limitBuilder[T]) Max(max int) *limitBuilder[T] {
	lb.params.max = max

	return lb
}

func (lb *limitBuilder[T]) Build() (t T) {
	switch target := any(lb.builder).(type) {
	case *builder:
		target.params.limit = lb.params
	case *addBuilder:
		target.self.limit = lb.params
	}
	t = lb.builder

	return
}
