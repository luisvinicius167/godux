/*
Package godux implements a state management for your backend application.
It's inspired in Redux, but with simplest concepts.
- State:   Your application state don't change.
- Actions: Your action is used in reducers, to return new value based on State.
- Reducers: Actions describe the fact that something happened, but don’t specify how the application’s state changes in response. This is the job of a reducer.

This library was inspired in Redux.
*/
package godux

import "sync"

// Store Your central store that has your application state
type Store struct {
	storeState
	storeStateLock sync.RWMutex
}

type storeState struct {
	state   map[string]interface{}
	reducer func(Action) interface{}
}

// Action that you create to change the State
type Action struct {
	Type  string
	Value interface{}
}

// NewStore to create your Store Application
func NewStore() *Store {
	return &Store{
		storeState: storeState{
			state: map[string]interface{}{},
		},
	}
}

// Reducer is a function that you use to return new value based on your storeState.
// Your state don't will be changed.
func (s *Store) Reducer(callback func(Action) interface{}) {
	s.storeStateLock.Lock()
	defer s.storeStateLock.Unlock()

	s.storeState.reducer = callback
}

// SetState is to sets the state store
func (s *Store) SetState(name string, value interface{}) {
	s.storeStateLock.Lock()
	defer s.storeStateLock.Unlock()

	s.storeState.state[name] = value
}

// Dispatch trigger your action type
func (s *Store) Dispatch(actionType Action) interface{} {
	s.storeStateLock.RLock()
	defer s.storeStateLock.RUnlock()	

	if s.storeState.reducer == nil {
		s.storeStateLock.RUnlock()
		panic("reducer not initialized")
	}
	ret := s.storeState.reducer(actionType)
	return ret
}

// GetState return the state of your store
func (s *Store) GetState(name string) interface{} {
	s.storeStateLock.RLock()
	defer s.storeStateLock.RUnlock()

	ret := s.storeState.state[name]
	return ret
}

// GetAllState return a full copy of the current state.
func (s *Store) GetAllState() interface{} {
	s.storeStateLock.RLock()
	defer s.storeStateLock.RUnlock()

	ret := map[string]interface{}{}	
	for k, v := range s.storeState.state {
		ret[k] = v
	}
	return ret
}
