package schedule_test

import (
	"testing"
	"time"

	"github.com/goexl/schedule"
)

func TestImmediately(t *testing.T) {
	_, _ = schedule.New().Build().Add(newImmediatelyWorker()).Build().Do()
	time.Sleep(time.Second)
	if 1 != immediately {
		t.Fatalf("期望：1，实际：%d", immediately)
	}
}
