package schedule

import (
	"github.com/goexl/log"
)

type params struct {
	logger log.Logger
	unique bool
	limit  *limitParams
}

func newParams() *params {
	return &params{
		logger: log.New().Apply(),
	}
}
