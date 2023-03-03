package schedule

import (
	"strings"
	"time"

	"github.com/goexl/gox"
)

type runtime struct {
	time time.Time
}

func newRuntime(time time.Time) *runtime {
	return &runtime{
		time: time,
	}
}

func (r *runtime) spec() string {
	builder := new(strings.Builder)
	// 秒
	builder.WriteString(gox.ToString(r.time.Second()))
	builder.WriteString(space)

	// 分钟
	builder.WriteString(gox.ToString(r.time.Minute()))
	builder.WriteString(space)

	// 小时
	builder.WriteString(gox.ToString(r.time.Hour()))
	builder.WriteString(space)

	// 日
	builder.WriteString(gox.ToString(r.time.Day()))
	builder.WriteString(space)

	// 月
	builder.WriteString(gox.ToString(r.time.Month()))
	builder.WriteString(space)

	// 周
	builder.WriteString(gox.ToString(r.time.Weekday()))
	builder.WriteString(space)

	return builder.String()
}
