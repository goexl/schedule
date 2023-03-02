package schedule

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/robfig/cron/v3"
)

var _ cron.Job = (*jobDefault)(nil)

type jobDefault struct {
	logger simaqian.Logger
	worker worker
}

func newDefaultJob(logger simaqian.Logger, worker worker) *jobDefault {
	return &jobDefault{
		logger: logger,
		worker: worker,
	}
}

func (jd *jobDefault) Run() {
	fields := gox.Fields[any]{
		field.New("worker", jd.worker),
	}
	if err := jd.worker.Run(); nil != err {
		jd.logger.Warn("任务执行出错", fields.Add(field.Error(err))...)
	} else {
		jd.logger.Debug("任务执行成功", fields...)
	}
}
