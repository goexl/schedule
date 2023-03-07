package schedule_test

import (
	"testing"
	"time"

	"github.com/goexl/schedule"
)

func TestImmediately(t *testing.T) {
	schedule.New().Build().Add(newImmediatelyWorker()).Build().Put()
	time.Sleep(time.Second)
	if 1 != immediately {
		t.Fatalf("期望：1，实际：%d", immediately)
	}
}
