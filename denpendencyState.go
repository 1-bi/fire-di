package fire_di

// dependencyState set the dependency state for proxy bean
type dependencyState struct {
	stateInjected  int8
	dependencyType string
}

func newDependencyState(objType string) *dependencyState {
	dpendencyState := new(dependencyState)
	dpendencyState.dependencyType = objType
	return dpendencyState
}

func (myself *dependencyState) updateState(newstate int8) {
	myself.stateInjected = newstate
}
