package schedule

type worker interface {
	// Run 执行任务
	Run() (err error)
}
