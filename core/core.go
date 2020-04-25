// Package core orchestrates every parts.
// It holds the domain logic.
package core

// STATE0 is the initial state.
const STATE0 = State{}

var (
	sim   = Simulation{}
	guide = Guide{sim}
	ai    = Agent{StateDimension: guide.EndDimension, ActionDimension: sim.ActionDimension}
	state = STATE0
)

// Restart sets state back to the inital state.
func Restart() {
	state = STATE0
}

// Step moves the simulation forward in time, and returns this next state.
func Step() State {
	action := ai.Act(guide.Parse(state))
	state = sim.Simulate(state, action)
	return state
}

// Train forces ai to get better at its job.
func Train() {
	ai.Train(guide)
}
