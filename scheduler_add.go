package schedule

import (
	"github.com/goexl/gox"
	"github.com/robfig/cron/v3"
)

func (s *Scheduler) add(params *addParams) (id string, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if eid, ae := s.addToCron(&id, params); nil != ae {
		err = ae
	} else {
		id = gox.Ift("" != params.id, params.id, gox.ToString(eid))
		s.ids.Store(id, eid)
	}

	return
}

func (s *Scheduler) addToCron(id *string, params *addParams) (eid cron.EntryID, err error) {
	// 检查任务是否已经存在
	// 判断逻辑，优先看添加参数里是否要求唯一
	// 再看整体配置里是否要求唯一
	unique := gox.Ift(params.unique, params.unique, gox.If(s.params.unique, s.params.unique))
	if unique && s.Contains(params.id) {
		return
	}

	// 检查限制是否符合要求，如果不符合，不返回任何值
	if err = params.checkLimit(s); nil != err {
		return
	}

	switch params.typ {
	case typeCron, typeDuration:
		job := newDefaultJob(params.worker, s.params.logger)
		tick := params.ticker.tick()
		eid, err = s.cron.AddJob(tick, job)
	case typeFixed:
		job := newOnceJob(id, s, params.worker, s.params.logger)
		eid = s.cron.Schedule(params.schedule, job)
	case typeImmediately:
		job := newOnceJob(id, s, params.worker, s.params.logger)
		eid = s.cron.Schedule(params.schedule, job)
	case typeRandom:
		job := newDefaultJob(params.worker, s.params.logger)
		eid = s.cron.Schedule(params.schedule, job)
	}

	return
}
