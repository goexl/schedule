package schedule

import (
	"sync"

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

func (s *Scheduler) Add(executor worker, opts ...option) (id string, err error) {

}

func (s *Scheduler) Start(_ ...option) {
	if !s.started {
		s.cron.Start()
		s.started = true
	}
}

func (s *Scheduler) Stop(_ ...option) {
	if !s.stopped {
		s.cron.Stop()
		s.stopped = true
	}
}

func (s *Scheduler) Remove(id *optionId) {
	if _id, ok := s.ids.Load(id.id); ok {
		s.cron.Remove(_id.(cron.EntryID))
		s.ids.Delete(id.id)
	}
}

func (s *Scheduler) Clear() {
	for _, entry := range s.cron.Entries() {
		s.cron.Remove(entry.ID)
	}
	s.ids = new(sync.Map)
}

func (s *Scheduler) Contains(id *optionId) (contains bool) {
	_, contains = s.ids.Load(id.id)

	return
}

func (s *Scheduler) Count() int {
	return len(s.cron.Entries())
}
