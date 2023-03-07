package schedule

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/robfig/cron/v3"
)

type jobOnce struct {
	id     *cron.EntryID
	cron   *cron.Cron
	worker   worker
	logger   simaqian.Logger
}

func newOnceJob(id *cron.EntryID, cron *cron.Cron, worker worker, logger simaqian.Logger) *jobOnce {
	return &jobOnce{
		id:id,
		cron:cron,
		worker:   worker,
		logger:   logger,
	}
}

func (jo *jobOnce) Run() {
	// 删除原来的任务，确保不会再被执行
	defer jo.cron.Remove(*jo.id)

	fields := gox.Fields[any]{
		field.New("worker", jo.worker),
	}
	if err := jo.worker.Run(); nil != err {
		jo.logger.Warn("任务执行出错", fields.Add(field.Error(err))...)
	} else {
		jo.logger.Debug("任务执行成功", fields...)
	}
}
