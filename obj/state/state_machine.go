package state

type StateMachine struct {
	CurrentState State
	States       map[string]State
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		States: make(map[string]State),
	}
}

func (sm *StateMachine) AddState(name string, state State) {
	sm.States[name] = state
}

func (sm *StateMachine) ChangeState(name string) {
	if newState, exists := sm.States[name]; exists {
		if sm.CurrentState != nil {
			sm.CurrentState.Exit()
		}
		sm.CurrentState = newState
		sm.CurrentState.Enter()
	}
}

func (sm *StateMachine) Update(dt float64) {
	if sm.CurrentState != nil {
		sm.CurrentState.Update(dt)
	}
}
