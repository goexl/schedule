package schedule

import (
	"time"

	"github.com/goexl/simaqian"
)

type params struct {
	logger simaqian.Logger
	unique bool
	echo   bool
	delay  time.Duration
}

func newParams() *params {
	return &params{
		logger: simaqian.Default(),
	}
}
