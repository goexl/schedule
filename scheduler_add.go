package schedule

import (
	"fmt"
	"strconv"

	"github.com/robfig/cron/v3"
	"github.com/rs/xid"
)

func (s *Scheduler) add(_options *addParams) (id string, err error) {
	if "" == _options.id {
		id = xid.New().String()
	}

	var entryId cron.EntryID
	switch _options.scheduleType {
	case typeCron:
		entryId, err = s.cron.AddFunc(_options.cron, func() {
			_ = executor.Run()
		})
	case typeDuration:
		entryId, err = s.cron.AddFunc(fmt.Sprintf("@every %s", _options.duration.String()), func() {
			_ = executor.Run()
		})
	case typeTime:
		entryId, err = s.cron.AddFunc(fixTimeSpec(_options.time, _options.delayMaxRand, _options.delay), func() {
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
