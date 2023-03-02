package schedule

type ticker interface {
	tick() string
}
