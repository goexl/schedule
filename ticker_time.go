package schedule

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var _ ticker = (*timeTicker)(nil)

type timeTicker struct {
	time time.Time
}

func newTimeTicker(time time.Time) *timeTicker {
	return &timeTicker{
		time: time,
	}
}

func (tt *timeTicker) tick() string {
	builder := new(strings.Builder)
	builder.WriteString(every)
	builder.WriteString(space)
	builder.WriteString(tt.time.String())

	return builder.String()
}

func (tt *timeTicker) fixTimeSpec(runtime time.Time, max int64, delay time.Duration) string {
	now := time.Now()
	if runtime.Before(now) {
		rand.Seed(time.Now().Unix())
		runtime = now.Add(time.Duration(rand.Int63n(max)) * time.Second)
	}
	if 0 != delay {
		runtime.Add(delay)
	}

	return fmt.Sprintf(
		"%d %d %d %d %d %d",
		runtime.Second(), runtime.Minute(), runtime.Hour(),
		runtime.Day(), runtime.Month(), runtime.Weekday(),
	)
}
