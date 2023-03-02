package schedule

import (
	"github.com/robfig/cron/v3"
)

var _ cron.Job = (*jobDefault)(nil)

type jobDefault struct {
	worker worker
}

func newDefaultJob(worker worker) *jobDefault {
	return &jobDefault{
		worker: worker,
	}
}

func (jd *jobDefault) Run() {

}
