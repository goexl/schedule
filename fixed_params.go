package schedule

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/goexl/gox"
)

type fixedParams struct {
	runtime time.Time
	random  bool
	min     time.Duration
	max     time.Duration
}

func newFixedParams(runtime time.Time) *fixedParams {
	return &fixedParams{
		runtime: runtime,
		random:  false,
	}
}

func (fp *fixedParams) delay() (delay time.Duration) {
	diff := fp.max - fp.min
	delay = gox.Ifx(fp.random, func() time.Duration {
		return fp.rand(diff)
	}, func() time.Duration {
		return diff
	})

	return
}

func (fp *fixedParams) rand(diff time.Duration) (delay time.Duration) {
	if duration, re := rand.Int(rand.Reader, big.NewInt(int64(diff))); nil != re {
		delay = time.Duration(duration.Int64())
	} else {
		delay = diff
	}

	return
}

func (fp *fixedParams) tick() string {
	now := time.Now()
	if fp.runtime.Before(now) {
		fp.runtime = now
	}

	delay := fp.delay()
	if 0 < delay {
		fp.runtime = fp.runtime.Add(delay)
	}

	// 如果当前时间还是比当前时间小，需要增加一点延迟保证任务一定会被调度
	if fp.runtime.Before(time.Now()) {
		fp.runtime = time.Now().Add(time.Second)
	}

	return newRuntime(fp.runtime).spec()
}
