package schedule

import (
	"reflect"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/robfig/cron/v3"
)

var _ cron.Job = (*jobDefault)(nil)

type jobDefault struct {
	worker worker
	logger log.Logger

	name string
}

func newDefaultJob(worker worker, logger log.Logger) *jobDefault {
	return &jobDefault{
		worker: worker,
		logger: logger,

		name: reflect.TypeOf(worker).String(),
	}
}

func (jd *jobDefault) Run() {
	fields := gox.Fields[any]{
		field.New("worker", jd.name),
	}
	if err := jd.worker.Run(); nil != err {
		jd.logger.Warn("任务执行出错", fields.Add(field.Error(err))...)
	} else {
		jd.logger.Debug("任务执行成功", fields...)
	}
}
