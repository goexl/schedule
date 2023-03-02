package schedule

type addParams struct {
	typ    typ
	worker worker
	ticker ticker
	id     string
}

func newAddParams(worker worker) *addParams {
	return &addParams{
		typ:    typeFixed,
		worker: worker,
	}
}
