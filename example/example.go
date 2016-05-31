package main

import (
	"fmt"

	"github.com/luisvinicius167/godux"
)

func main() {
	// Creating new Store
	store := godux.NewStore()
	store.SetState("count", 1)

	// Creating new Action

	increment := func(number int) godux.Action {
		return godux.Action{
			Type:  "INCREMENT",
			Value: number,
		}
	}

	decrement := func(number int) godux.Action {
		return godux.Action{
			Type:  "DECREMENT",
			Value: number,
		}
	}
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
	// Receive new value
	newCount := store.Dispatch(increment(10))
	subCount := store.Dispatch(decrement(10))
	fmt.Printf("Your Store state is: %s. Your newCount is: %s. Your subCount is: %s", store.GetState("count"), newCount, subCount)
}
