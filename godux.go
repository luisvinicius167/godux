package godux

// Store has your applicatio state
type Store struct {
}

type storeState struct {
	state    map[string]interface{}
	reductor func(acton string) interface{}
}

// Action in Godux
type Action struct {
	Type map[string]string
}

var storestate *storeState

//NewStore create new Store
func NewStore() *Store {
	storestate = &storeState{}
	return new(Store)
}

// Reductor function to reduce
func (s *Store) Reductor(callback func(actionName string) interface{}) {
	storestate.reductor = callback
}

// Setstate are setting the state store
func (s *Store) Setstate(name string, value interface{}) {
	if len(storestate.state) == 0 {
		storestate.state = make(map[string]interface{})
	}
	storestate.state[name] = value
}

// Dispatch trigger for state change
func (s *Store) Dispatch(actionType string) interface{} {
	return storestate.reductor(actionType)
}

// CreateAction Create an action with the action type name
func (action *Action) CreateAction(names ...string) {
	if len(action.Type) == 0 {
		action.Type = make(map[string]string)
	}
	for _, name := range names {
		action.Type[name] = name
	}
}

// GetState return the state
func (s *Store) GetState(name string) interface{} {
	return storestate.state[name]
}

// GetAllState return the state
func (s *Store) GetAllState() interface{} {
	return storestate.state
}
