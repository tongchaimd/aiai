// Package core orchestrates every parts.
// It holds the domain logic.
package core

type State interface{}
type Action interface{}

type simulation interface {
	STATE0() State
	Simulate(current State, a Action) (next State)
}

type agent interface {
	Train()
	Act(State) Action
}

type agentConstructor interface {
	Agent(simulation) agent
}

// Core struct holds some fields, but is not indended to be created by literal.
type Core struct {
	sim       simulation
	ai        agent
	currState State
}

// Core function is the constructor to core.
// It initializes state to zero state, simulation implimentation, and agent implementation.
func NewCore(simImpl simulation, agentCon agentConstructor) (obj Core) {
	obj = Core{
		sim:       simImpl,
		ai:        agentCon.Agent(simImpl),
		currState: simImpl.STATE0(),
	}
	return
}

// Restart sets state back to the inital state.
func (c *Core) Restart() {
	c.currState = nil
}

// Step moves the simulation forward in time, and returns this next state.
func (c *Core) Step() State {
	action := c.ai.Act(c.currState)
	c.currState = c.sim.Simulate(c.currState, action)
	return c.currState
}

// Train forces ai to get better at its job.
func (c *Core) Train() {
	c.ai.Train()
}
