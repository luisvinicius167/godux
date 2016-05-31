/*
Package godux implements a state management for your backend application.
It's inspired in Redux, but with simplest concepts.
- State:   Your application state don't change.
- Actions: Your action is used in reducers, to return new value based on State.
- Reducers: Actions describe the fact that something happened, but don’t specify how the application’s state changes in response. This is the job of a reducer.

This library was inspired in Redux.
*/
package godux

// Store Your central store that has your application state
type Store struct{}

type storeState struct {
	state   map[string]interface{}
	reducer func(action Action) interface{}
}

// Action that you create to change the State
type Action struct {
	Type  string
	Value interface{}
}

var storestate *storeState

// NewStore to create your Store Application
func NewStore() *Store {
	storestate = &storeState{}
	return new(Store)
}

// Reducer is a function that you use to return new value based on your storeState.
// Your state don't will be changed.
func (s *Store) Reducer(callback func(action Action) interface{}) {
	storestate.reducer = callback
}

// Setstate is to sets the state store
func (s *Store) Setstate(name string, value interface{}) {
	if len(storestate.state) == 0 {
		storestate.state = make(map[string]interface{})
	}
	storestate.state[name] = value
}

// Dispatch trigger your action type
func (s *Store) Dispatch(actionType Action) interface{} {
	return storestate.reducer(actionType)
}

// GetState return the state of your store
func (s *Store) GetState(name string) interface{} {
	return storestate.state[name]
}

// GetAllState return all store state
func (s *Store) GetAllState() interface{} {
	return storestate.state
}
