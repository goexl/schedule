package schedule

import (
	"math"
	"time"

	"github.com/goexl/gox"
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
	return lp.checkCount(scheduler) && lp.checkProcess() && lp.checkCpu() && lp.checkMemory()
}

func (lp *limitParams) checkMemory() (success bool) {
	vms, vme := mem.VirtualMemory()
	success = gox.If(nil == vme, vms.UsedPercent < lp.memory)

	return
}

func (lp *limitParams) checkCpu() (success bool) {
	percent, pe := cpu.Percent(time.Second, false)
	success = gox.If(nil == pe, percent[0] < lp.cpu)

	return
}

func (lp *limitParams) checkProcess() (success bool) {
	ids, pe := process.Pids()
	success = gox.If(nil == pe, len(ids) <= lp.process)

	return
}

func (lp *limitParams) checkCount(scheduler *Scheduler) bool {
	return scheduler.Count() <= lp.max
}
