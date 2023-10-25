package schedule

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
)

type jobOnce struct {
	id        *string
	scheduler *Scheduler
	worker    worker

	logger   log.Logger
	deleted  bool
	executed bool
}

func newOnceJob(id *string, scheduler *Scheduler, worker worker, logger log.Logger) *jobOnce {
	return &jobOnce{
		id:        id,
		scheduler: scheduler,
		worker:    worker,

		logger: logger,
	}
}

func (jo *jobOnce) Run() {
	// 在任何情况下，确保任务被删除
	defer jo.cleanup()

	// 只能被执行一次
	if jo.executed {
		return
	}

	jo.executed = true
	fields := gox.Fields[any]{
		field.New("worker", jo.worker),
	}
	if err := jo.worker.Run(); nil != err {
		jo.logger.Warn("任务执行出错", fields.Add(field.Error(err))...)
	} else {
		jo.logger.Debug("任务执行成功", fields...)
	}

	// 删除原来的任务，确保不会再被执行
	jo.scheduler.remove(*jo.id)
	jo.deleted = true
}

func (jo *jobOnce) cleanup() {
	if !jo.deleted {
		jo.scheduler.remove(*jo.id)
	}
}
