package schedule

import (
	"github.com/goexl/gox"
)

type removeBuilder struct {
	scheduler *Scheduler
	params    *removeParams
}

func newRemoveBuilder(scheduler *Scheduler) *removeBuilder {
	return &removeBuilder{
		scheduler: scheduler,
		params:    newRemoveParams(),
	}
}

func (rb *removeBuilder) Id(ids ...any) *removeBuilder {
	for _, _id := range ids {
		switch target := _id.(type) {
		case id:
			rb.params.filters = append(rb.params.filters, filterEqual(target.TaskId()))
		default:
			rb.params.filters = append(rb.params.filters, filterEqual(gox.ToString(target)))
		}
	}

	return rb
}

func (rb *removeBuilder) Filter(filters ...filter) *removeBuilder {
	rb.params.filters = append(rb.params.filters, filters...)

	return rb
}

func (rb *removeBuilder) Build() *remove {
	return newRemove(rb.scheduler, rb.params)
}
