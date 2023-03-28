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

func (s *Scheduler) Remove(ids ...any) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, _id := range ids {
		taskId := ""
		switch target := _id.(type) {
		case id:
			taskId = target.Id()
		default:
			taskId = gox.ToString(target)
		}

		if entry, ok := s.ids.Load(taskId); ok {
			s.cron.Remove(entry.(cron.EntryID))
			s.ids.Delete(taskId)
		}
	}
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
