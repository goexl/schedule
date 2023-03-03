package schedule

import (
	"time"
)

var _ ticker = (*immediatelyTicker)(nil)

type immediatelyTicker struct{}

func newImmediatelyTicker() *immediatelyTicker {
	return new(immediatelyTicker)
}

func (it *immediatelyTicker) tick() string {
	return newRuntime(time.Now().Add(100 * time.Millisecond)).spec()
}
