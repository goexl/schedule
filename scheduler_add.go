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
	switch params.typ {
	case typeCron, typeDuration, typeFixed:
		id, err = s.cron.AddJob(params.ticker.tick(), newDefaultJob(s.logger, params.worker))
	}

	return
}
