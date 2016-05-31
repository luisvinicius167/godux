/*
Package godux implements a state management for your backend application.
It's inspired in Redux, but with simplest concepts.
- State:   Your application state don't change.
- Actions: Your action is used in reducers, to return new value based on State.

This library was inspired in Redux.
*/
package godux

// Store Your central store that has your application state
type Store struct{}

type storeState struct {
	state    map[string]interface{}
	reductor func(acton string) interface{}
}

// Action that you create to change the State
type Action struct {
	Type map[string]string
}

var storestate *storeState

// NewStore to create your Store Application
func NewStore() *Store {
	storestate = &storeState{}
	return new(Store)
}

// Reductor is a function that you use to return new value based on your storeState.
// Your state don't will be changed.
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

// Dispatch trigger your action type
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

// GetState return the state of your store
func (s *Store) GetState(name string) interface{} {
	return storestate.state[name]
}

// GetAllState return all store state
func (s *Store) GetAllState() interface{} {
	return storestate.state
}
