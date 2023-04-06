package schedule

import (
	"github.com/goexl/simaqian"
)

type params struct {
	logger simaqian.Logger
	unique bool
	limit  *limitParams
}

func newParams() *params {
	return &params{
		logger: simaqian.Default(),
	}
}
