package schedule

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

type jobOnce struct {
	schedule *scheduleOnce
	worker   worker
	logger   simaqian.Logger
}

func newOnceJob(schedule *scheduleOnce, worker worker, logger simaqian.Logger) *jobOnce {
	return &jobOnce{
		schedule: schedule,
		worker:   worker,
		logger:   logger,
	}
}

func (jo *jobOnce) Run() {
	// 必须执行完成
	defer jo.schedule.completed()

	fields := gox.Fields[any]{
		field.New("worker", jo.worker),
	}
	if err := jo.worker.Run(); nil != err {
		jo.logger.Warn("任务执行出错", fields.Add(field.Error(err))...)
	} else {
		jo.logger.Debug("任务执行成功", fields...)
	}
}
