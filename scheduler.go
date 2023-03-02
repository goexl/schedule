package schedule

import (
	"sync"

	"github.com/goexl/gox"
	"github.com/robfig/cron/v3"
)

// Scheduler 任务计划组织程序
type Scheduler struct {
	cron *cron.Cron
	ids  *sync.Map

	started bool
	stopped bool
}

func newScheduler() (scheduler *Scheduler) {
	scheduler = &Scheduler{
		cron: cron.New(cron.WithSeconds()),
		ids:  new(sync.Map),

		started: false,
		stopped: false,
	}
	scheduler.Start()

	return
}

func (s *Scheduler) Add(worker worker) *addBuilder {
	return newAddBuilder(s, worker)
}

func (s *Scheduler) Start() {
	if !s.started {
		s.cron.Start()
		s.started = true
	}
}

func (s *Scheduler) Stop() {
	if !s.stopped {
		s.cron.Stop()
		s.stopped = true
	}
}

func (s *Scheduler) Remove(id any) {
	id = gox.ToString(id)
	if entry, ok := s.ids.Load(id); ok {
		s.cron.Remove(entry.(cron.EntryID))
		s.ids.Delete(id)
	}
}

func (s *Scheduler) Clear() {
	for _, entry := range s.cron.Entries() {
		s.cron.Remove(entry.ID)
	}
	s.ids = new(sync.Map)
}

func (s *Scheduler) Contains(id any) (contains bool) {
	_, contains = s.ids.Load(gox.ToString(id))

	return
}

func (s *Scheduler) Count() int {
	return len(s.cron.Entries())
}
