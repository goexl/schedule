package schedule

var _ ticker = (*cronTicker)(nil)

type cronTicker struct {
	cron string
}

func newCronTicker(cron string) *cronTicker {
	return &cronTicker{
		cron: cron,
	}
}

func (ct *cronTicker) tick() string {
	return ct.cron
}
