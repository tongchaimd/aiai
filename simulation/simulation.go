package simulation

// Obj is a dummy struct to workaround interface stuff.
type Obj struct{}

type State struct{}
type Action struct{}

func (o Obj) Simulate(current State, action Action) (next State) {
	next = State{}
	return
}

func (o Obj) STATE0() State {
	return State{}
}
