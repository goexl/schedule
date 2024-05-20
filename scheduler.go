package schedule

import (
	"sync"

	"github.com/goexl/gox"
	"github.com/robfig/cron/v3"
)

// Scheduler 任务计划组织程序
type Scheduler struct {
	cron   *cron.Cron
	ids    *sync.Map
	params *params

	mutex   *sync.Mutex
	started bool
	stopped bool
}

func newScheduler(params *params) (scheduler *Scheduler) {
	scheduler = &Scheduler{
		params: params,
		cron:   cron.New(cron.WithSeconds()),
		ids:    new(sync.Map),

		mutex:   new(sync.Mutex),
		started: false,
		stopped: false,
	}
	scheduler.Start()

	return
}

func (s *Scheduler) Add(worker worker) *addBuilder {
	return newAddBuilder(s, worker, s.params)
}

func (s *Scheduler) Start() (err error) {
	if !s.started {
		s.cron.Start()
		s.started = true
	}

	return
}

func (s *Scheduler) Stop() (err error) {
	if !s.stopped {
		s.cron.Stop()
		s.stopped = true
	}

	return
}

func (s *Scheduler) Remove() *removeBuilder {
	return newRemoveBuilder(s)
}

func (s *Scheduler) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, entry := range s.cron.Entries() {
		s.cron.Remove(entry.ID)
	}
	s.ids = new(sync.Map)
}

func (s *Scheduler) Contains(id any) (contains bool) {
	id = gox.ToString(id)
	_, contains = s.ids.Load(id)

	return
}

func (s *Scheduler) Count() int {
	return len(s.cron.Entries())
}

func (s *Scheduler) remove(id string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if entry, ok := s.ids.Load(id); ok {
		s.cron.Remove(entry.(cron.EntryID))
		s.ids.Delete(id)
	}
}

func (s *Scheduler) checkLimit() (err error) {
	if nil != s.params.limit {
		err = s.params.limit.check(s)
	}

	return
}
