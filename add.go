package schedule

type add struct {
	scheduler *Scheduler
	params    *addParams
}

func newAdd(scheduler *Scheduler, params *addParams) *add {
	return &add{
		scheduler: scheduler,
		params:    params,
	}
}

func (a *add) Apply() (string, error) {
	return a.scheduler.add(a.params)
}
