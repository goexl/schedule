package schedule

import (
	"github.com/goexl/simaqian"
)

type params struct {
	logger simaqian.Logger
	unique *bool
}

func newParams() *params {
	return &params{
		logger: simaqian.Default(),
	}
}
