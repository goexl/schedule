package schedule

type remove struct {
	scheduler *Scheduler
	params    *removeParams
}

func newRemove(scheduler *Scheduler, params *removeParams) *remove {
	return &remove{
		scheduler: scheduler,
		params:    params,
	}
}

func (r *remove) Do() {
	r.scheduler.ids.Range(func(key, _ any) bool {
		if r.check(key.(string)) {
			r.scheduler.remove(key.(string))
		}

		return true
	})
}

func (r *remove) check(id string) (checked bool) {
	for _, _filter := range r.params.filters {
		checked = _filter(id)
		if checked {
			break
		}
	}

	return
}
