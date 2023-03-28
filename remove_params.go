package schedule

type removeParams struct {
	filters []filter
}

func newRemoveParams() *removeParams {
	return new(removeParams)
}
