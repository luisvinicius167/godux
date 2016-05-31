package godux

import "testing"

func TestStoreStateCount(t *testing.T) {
	store := NewStore()
	store.Setstate("count", 1)
	if store.GetState("count") != 1 {
		t.Error("Expected 1 for Count State value.")
	}

}
func TestActionType(t *testing.T) {
	store := NewStore()
	store.Setstate("count", 1)

	increment := func(number int) Action {
		return Action{
			Type:  "INCREMENT",
			Value: number,
		}
	}

	reducer := func(action Action) interface{} {
		switch action.Type {
		case "INCREMENT":
			return store.GetState("count").(int) + action.Value.(int)
		default:
			return store.GetAllState()
		}
	}
	store.Reducer(reducer)
	value := store.Dispatch(increment(1))
	if value != 2 {
		t.Error("The value are different. Action not correct.")
	}
}
