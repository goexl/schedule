package schedule

const (
	typeCron typ = iota
	typeDuration
	typeFixed
	typeImmediately
	typeRandom
)

type typ uint8
