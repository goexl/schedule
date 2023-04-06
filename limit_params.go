package schedule

import (
	"math"
	"time"

	"github.com/goexl/exc"
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
		err = exc.NewFields("内存限制不通过", field.New("memory", vms.UsedPercent), field.New("limit", lp.memory))
	}

	return
}

func (lp *limitParams) checkCpu() (err error) {
	if percent, pe := cpu.Percent(time.Second, false); nil != pe {
		err = pe
	} else if percent[0] >= lp.cpu {
		err = exc.NewFields("Cpu限制不通过", field.New("cpu", percent[0]), field.New("limit", lp.cpu))
	}

	return
}

func (lp *limitParams) checkProcess() (err error) {
	if ids, pe := process.Pids(); nil != pe {
		err = pe
	} else if len(ids) >= lp.process {
		err = exc.NewFields("进程数量限制不通过", field.New("process", len(ids)), field.New("limit", lp.process))
	}

	return
}

func (lp *limitParams) checkCount(scheduler *Scheduler) (err error) {
	count := scheduler.Count()
	if count > lp.max {
		err = exc.NewFields("任务数量限制不通过", field.New("count", count), field.New("limit", lp.max))
	}

	return
}
