package schedule

const (
	typeCron typ = iota
	typeDuration
	typeFixed
	typeImmediately
	typeRandomDuration
	typeRandomTime
)

type typ uint8
