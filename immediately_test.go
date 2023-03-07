package schedule_test

var immediately = 0

type immediatelyWorker struct{}

func newImmediatelyWorker() *immediatelyWorker {
	return new(immediatelyWorker)
}

func (iw *immediatelyWorker) Run() (err error) {
	immediately++

	return
}
