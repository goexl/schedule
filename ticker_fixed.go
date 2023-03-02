package schedule

var _ ticker = (*fixedTicker)(nil)

type fixedTicker struct {
	params *fixedParams
}

func newFixedTicker(params *fixedParams) *fixedTicker {
	return &fixedTicker{
		params: params,
	}
}

func (ft *fixedTicker) tick() string {
	return ft.params.tick()
}
