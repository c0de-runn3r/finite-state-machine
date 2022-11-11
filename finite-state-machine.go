package main

type State struct {
	Name string
}

type StateMachine struct {
	CurrentState   State
	ExistingStates []State
}

func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

func (sm *StateMachine) NewState(StateName string) *State {
	state := State{Name: StateName}
	sm.ExistingStates = append(sm.ExistingStates, state)
	return &state
}

func (sm *StateMachine) SetState(state State) {
	sm.CurrentState = state
}
