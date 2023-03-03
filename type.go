package schedule

const (
	typeCron typ = iota
	typeDuration
	typeFixed
	typeImmediately
)

type typ uint8
