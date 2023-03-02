package schedule

import (
	"github.com/goexl/gox"
	"github.com/robfig/cron/v3"
	"github.com/rs/xid"
)

func (s *Scheduler) add(params *addParams) (id string, err error) {
	if eid, ae := s.addToCron(params); nil == ae {
		id = gox.Ift("" == params.id, xid.New().String(), gox.ToString(eid))
		s.ids.Store(id, eid)
	}

	return
}

func (s *Scheduler) addToCron(params *addParams) (id cron.EntryID, err error) {
	// 检查任务是否已经存在
	if s.params.unique && s.Contains(params.id) {
		return
	}

	// 检查限制是否符合要求，如果不符合，不返回任何值
	if !params.limit.check(s) {
		return
	}

	switch params.typ {
	case typeCron, typeDuration, typeFixed:
		id, err = s.cron.AddJob(params.ticker.tick(), newDefaultJob(s.params.logger, params.worker))
	}

	return
}
