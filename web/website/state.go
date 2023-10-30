package website

type State struct {
	IsLoggedIn bool
}

func BuildState() State {
	return State{
		IsLoggedIn: true,
	}
}
