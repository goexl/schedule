package schedule

func filterEqual(sid string) filter {
	return func(id string) bool {
		return id == sid
	}
}
