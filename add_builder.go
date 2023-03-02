package schedule

import (
	"github.com/goexl/gox"
)

type addBuilder struct {
	params *addParams
}

func newAddBuilder(worker worker) *addBuilder {
	return &addBuilder{
		params: newAddParams(worker),
	}
}

func (ab *addBuilder) Id(id any) *addBuilder {
	ab.params.id = gox.ToString(id)

	return ab
}
