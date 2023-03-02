package schedule

const (
	typeCron typ = iota
	typeDuration
	typeTime
)

type typ uint8
