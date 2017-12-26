# godux <br/>

[![Join the chat at https://gitter.im/luisvinicius167/godux](https://badges.gitter.im/luisvinicius167/godux.svg)](https://gitter.im/luisvinicius167/godux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/luisvinicius167/godux)](https://goreportcard.com/report/github.com/luisvinicius167/godux)
> State Management for Go Backend applications inspired by Redux.

<p align="center">
  <img src="img/godux_.png" alt="Godux">
</p>

<pre align="center">
╔═════════╗       ╔══════════╗       ╔═══════════╗       ╔═════════════════╗
║ Action  ║──────>║ Reducer  ║ ────> ║   Store   ║ ────> ║   Application   ║
╚═════════╝       ╚══════════╝       ╚═══════════╝       ╚═════════════════╝
     ^                                                            │
     └────────────────────────────────────────────────────────────┘

</pre>

### Install
* Go: ``` go get github.com/luisvinicius167/godux ```

### Data Flow
**godux** gives go unidirectional data flow:

* The Action returns a small map with specific directions that are dispatched to a Reducer.
* The Reducer is a pure function (pure functions don't change original arguments) if relevant to it returns a new Value.
* The Value becomes the new State of the Store.

### Principles:
* Global application state is held in the Store, as a single map.
* State is ready-only (only *change* it only by *replacing* it with the Reducer).
* Changes are made with pure functions - Actions/Reducers that do not change the actual object but make a changed copy.

### Store:
A Store is basically a container that holds your application state.

```go
    store := godux.NewStore()
	store.Setstate("count", 1)
    store.Setstate("Title", "I like godux!")
```

#### Action
Actions are just pure functions which pass on their inputs when they're dispatched. Actions are stored on the `godux` map as `godux.Action`. 

```go
    increment := func(number int) godux.Action {
		return godux.Action{
			Type:  "INCREMENT",
			Value: number,
		}
	}
```
### Reducers
As in Redux: 
> "Actions describe the fact that something happened, but don’t specify how the application’s state changes in response. This is the job of a reducer". 

Reducers are *pure* functions that take in actions and the state of the store as inputs and leave them all as they came in (aka. pure)-- especially the original state of the store must not be modified (it's accessed by `store.GetState`)).

```go
    // reducer function
	reducer := func(action godux.Action) interface{} {
		switch action.Type {
		case "INCREMENT":
			return store.GetState("count").(int) + action.Value.(int)
		case "DECREMENT":
			return action.Value.(int) - store.GetState("count").(int)
		default:
			return store.GetAllState()
		}
	}
	// Add your reducer function to return new values basend on your state
	store.Reducer(reducer)
```
#### Dispatch
Dispatching an action is very easy.
```go
    // Receive new value
	newCount := store.Dispatch(increment(1)) // return 2
```

### API Reference

* #### Store:
  * ` godux.newStore() `: Create a single store with the state of your application (should only be used once).
  * ` godux.SetState(name string, value interface{}) `: Sets the state of the store.
  * ` godux.GetState(name string) `: Return a state's value.
  * ` godux.GetAllState() `: Return the whole state as a map.

* ##### Reducer:
  * ``` store.Reducer(func(action godux.Action)) ```: Adding a reducer function to your Store.

* ##### Dispatch:
  * ``` store.Dispatch(action godux.Action) ```: Dispatching an action to your Reducer.

* #### Action:
  * ``` godux.Action( Type string, Value interface{}) ```: Adding an easily available Action.

### License
MIT License.
