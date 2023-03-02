package schedule

import (
	"fmt"
	"strconv"

	"github.com/robfig/cron/v3"
	"github.com/rs/xid"
)

func (s *Scheduler) add(params *addParams) (id string, err error) {
	if "" == params.id {
		id = xid.New().String()
	}

	var entryId cron.EntryID
	switch params.scheduleType {
	case typeCron:
		entryId, err = s.cron.AddFunc(params.cron, func() {
			_ = executor.Run()
		})
	case typeDuration:
		entryId, err = s.cron.AddFunc(fmt.Sprintf("@every %s", params.duration.String()), func() {
			_ = executor.Run()
		})
	case typeTime:
		entryId, err = s.cron.AddFunc(fixTimeSpec(params.time, params.delayMaxRand, params.delay), func() {
			_ = executor.Run()
		})
	}
	if nil != err {
		return
	}
	id = strconv.Itoa(int(entryId))
	s.ids.Store(id, entryId)

	return
}
