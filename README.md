# godux <br/>
> State Management for Go Backend applications inspired in Redux.

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
**godux** turns your data flow unidirectional:

* Create actions, like pure functions.
* Store dispatch the action.
* Return new Value based on your Store State.

### Principles:
* Application State is held in the store, as a single map.
* State are ready-only
* Changes are made with pure functions.

### Store:
A Store is basically a container that holds your application state.

```go
    store := godux.NewStore()
	store.Setstate("count", 1)
    store.Setstate("Title", "I like godux!")
```

#### Action
Actions are just pure function. Your Actions functions always return a godux.Action. 

```go
    increment := func(number int) godux.Action {
		return godux.Action{
			Type:  "INCREMENT",
			Value: number,
		}
	}
```
### Reducers
Like Redux Concept: "Actions describe the fact that something happened, but don’t specify how the application’s state changes in response. This is the job of a reducer."
```go
    // reductor function
	reductor := func(action godux.Action) interface{} {
		switch action.Type {
		case "INCREMENT":
			return store.GetState("count").(int) + action.Value.(int)
		case "DECREMENT":
			return action.Value.(int) - store.GetState("count").(int)
		default:
			return store.GetAllState()
		}
	}
	// Add your reductor function to return new values basend on your state
	store.Reductor(reductor)
```
#### Dispatch
Dispatch the an action is very easy.
```go
    // Receive new value
	newCount := store.Dispatch(increment(1)) // return 2
```

### API Reference

* #### Store:
  * `` godux.newStore() ```: Create a single store with the state of your application.
  * `` godux.SetState(name string, value interface{}) ```: Sets the state store.
  * `` godux.GetState(name string) ```: Return your Store state value.
  * `` godux.GetAllState() ```: Return all state value like a map type.

* #### Store Reductor:
  * ``` store.Reductor(func(action godux.Action)) ```: Add the reductor function in your Store.

* #### Store Dispatch:
  * ``` store.Dispatch(action godux.Action) ```: Dispatch your action to your Reducer.

* #### Action:
  * ``` godux.Action( Type string, Value interface{}) ```: Your application action.

### License
MIT License.
