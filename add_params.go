package schedule

type addParams struct {
	typ    typ
	worker worker
	ticker ticker
	limit  *limitParams
	id     string
	unique bool
}

func newAddParams(worker worker) *addParams {
	return &addParams{
		typ:    typeImmediately,
		worker: worker,
	}
}

func (ap *addParams) checkLimit(scheduler *Scheduler) (checked bool) {
	if nil == ap.limit {
		checked = true
	} else {
		checked = ap.limit.check(scheduler)
	}

	return
}
