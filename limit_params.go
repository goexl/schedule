package schedule

import (
	"math"
	"time"

	"github.com/goexl/exception"
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

func newUnlimitedParams() *limitParams {
	return &limitParams{
		cpu:     math.MaxFloat64,
		memory:  math.MaxFloat64,
		process: math.MaxInt,
		max:     math.MaxInt,
	}
}

func (lp *limitParams) check(scheduler *Scheduler) (err error) {
	if ce := lp.checkCount(scheduler); nil != ce {
		err = ce
	} else if pe := lp.checkProcess(); nil != pe {
		err = pe
	} else if ce := lp.checkCpu(); nil != ce {
		err = ce
	} else if me := lp.checkMemory(); nil != me {
		err = me
	}

	return
}

func (lp *limitParams) checkMemory() (err error) {
	if vms, vme := mem.VirtualMemory(); nil != vme {
		err = vme
	} else if vms.UsedPercent >= lp.memory {
		message := "内存限制不通过"
		memory := field.New("memory", vms.UsedPercent)
		limit := field.New("limit", lp.memory)
		err = exception.New().Message(message).Field(memory, limit).Build()
	}

	return
}

func (lp *limitParams) checkCpu() (err error) {
	if percent, pe := cpu.Percent(time.Second, false); nil != pe {
		err = pe
	} else if percent[0] >= lp.cpu {
		message := "Cpu限制不通过"
		_cpu := field.New("cpu", percent[0])
		limit := field.New("limit", lp.cpu)
		err = exception.New().Message(message).Field(_cpu, limit).Build()
	}

	return
}

func (lp *limitParams) checkProcess() (err error) {
	if ids, pe := process.Pids(); nil != pe {
		err = pe
	} else if len(ids) >= lp.process {
		message := "进程数量限制不通过"
		_process := field.New("process", lp.process)
		limit := field.New("limit", lp.process)
		err = exception.New().Message(message).Field(_process, limit).Build()
	}

	return
}

func (lp *limitParams) checkCount(scheduler *Scheduler) (err error) {
	count := scheduler.Count()
	if count > lp.max {
		message := "任务数量限制不通过"
		_count := field.New("count", count)
		limit := field.New("limit", lp.max)
		err = exception.New().Message(message).Field(_count, limit).Build()
	}

	return
}
