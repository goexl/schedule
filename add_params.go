package schedule

type addParams struct {
	typ    typ
	worker worker
	ticker ticker
	limit  *limitParams
	id     string
	unique *bool
}

func newAddParams(worker worker) *addParams {
	return &addParams{
		typ:    typeFixed,
		worker: worker,
	}
}
