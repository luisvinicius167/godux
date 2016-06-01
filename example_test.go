package godux_test

import (
	"fmt"

	"github.com/luisvinicius167/godux"
)

func Example() {
	// Creating new Store
	store := godux.NewStore()
	// Set state
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
			return store.GetState("count").(int) - action.Value.(int)
		default:
			return store.GetAllState()
		}
	}
	// Add your reducer function to return new values based on your state
	store.Reducer(reducer)
	// Receive new value
	newCount := store.Dispatch(increment(10)) // 1+10=11
	subCount := store.Dispatch(decrement(10)) // 1-10=-9
	fmt.Printf("Your Store state is: %d. Your newCount is: %d. Your subCount is: %d\n", store.GetState("count"), newCount, subCount)
	// output: Your Store state is: 1. Your newCount is: 11. Your subCount is: -9
}
