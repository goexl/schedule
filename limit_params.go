package schedule

import (
	"math"
)

type limitParams struct {
	cpu     int
	memory  int
	process int
}

func newLimitParams() *limitParams {
	return &limitParams{
		cpu:     100,
		memory:  100,
		process: math.MaxInt,
	}
}

func (lp *limitParams) check(scheduler *Scheduler) bool {
	return true
}
