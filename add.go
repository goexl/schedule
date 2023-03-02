package schedule

type add struct {
	params *addParams
}

func newAdd(params *addParams) *add {
	return &add{
		params: params,
	}
}

func (a *add) Put() (id string, err error) {

}
