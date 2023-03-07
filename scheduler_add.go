package schedule

import (
	"github.com/goexl/gox"
	"github.com/robfig/cron/v3"
)

func (s *Scheduler) add(params *addParams) (id string, err error) {
	if eid, ae := s.addToCron(params); nil == ae {
		id = gox.Ift("" != params.id, params.id, gox.ToString(eid))
		s.ids.Store(id, eid)
	}

	return
}

func (s *Scheduler) addToCron(params *addParams) (id cron.EntryID, err error) {
	// 检查任务是否已经存在
	// 判断逻辑，优先看添加参数里是否要求唯一
	// 再看整体配置里是否要求唯一
	unique := gox.Ift(params.unique, params.unique, gox.If(s.params.unique, s.params.unique))
	if unique && s.Contains(params.id) {
		return
	}

	// 检查限制是否符合要求，如果不符合，不返回任何值
	if !params.checkLimit(s) {
		return
	}

	switch params.typ {
	case typeCron, typeDuration, typeFixed:
		id, err = s.cron.AddJob(params.ticker.tick(), newDefaultJob(s.params.logger, params.worker))
	case typeImmediately:
		id = s.cron.Schedule(newScheduleOnce(&id, s.cron), newDefaultJob(s.params.logger, params.worker))
	}

	return
}
