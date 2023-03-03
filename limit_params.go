package schedule

import (
	"math"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type limitParams struct {
	cpu     float64
	memory  float64
	process int
	max     int
}

func newLimitParams() *limitParams {
	return &limitParams{
		cpu:     100,
		memory:  100,
		process: math.MaxInt,
		max:     math.MaxInt,
	}
}

func (lp *limitParams) check(scheduler *Scheduler) bool {
	return lp.checkCount(scheduler) && lp.checkProcess(scheduler) && lp.checkCpu(scheduler) && lp.checkMemory(scheduler)
}

func (lp *limitParams) checkMemory(scheduler *Scheduler) (success bool) {
	vms, vme := mem.VirtualMemory()
	success = gox.If(nil == vme, vms.UsedPercent < lp.memory)
	if !success {
		lp.log(scheduler, field.New("memory", vms.UsedPercent), field.New("limit", lp.memory))
	}

	return
}

func (lp *limitParams) checkCpu(scheduler *Scheduler) (success bool) {
	percent, pe := cpu.Percent(time.Second, false)
	success = gox.If(nil == pe, percent[0] < lp.cpu)
	if !success {
		lp.log(scheduler, field.New("cpu", percent[0]), field.New("limit", lp.cpu))
	}

	return
}

func (lp *limitParams) checkProcess(scheduler *Scheduler) (success bool) {
	ids, pe := process.Pids()
	process := len(ids)
	success = gox.If(nil == pe, process <= lp.process)
	if !success {
		lp.log(scheduler, field.New("process", process), field.New("limit", lp.process))
	}

	return
}

func (lp *limitParams) checkCount(scheduler *Scheduler) (success bool) {
	count := scheduler.Count()
	success = count <= lp.max
	if !success {
		lp.log(scheduler, field.New("count", count), field.New("limit", lp.max))
	}

	return
}

func (lp *limitParams) log(scheduler *Scheduler, fields ...gox.Field[any]) {
	scheduler.params.logger.Info("限制检查未通过", fields...)
}
