package schedule

const (
	typeCron typ = iota
	typeDuration
	typeFixed
)

type typ uint8
