package schedule

var _ = New

// New 创建计划任务
func New() *builder {
	return newBuilder()
}
