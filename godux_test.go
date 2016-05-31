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
	action := Action{}
	action.CreateAction("INCREMENT")
	reductor := func(actionType string) interface{} {
		switch actionType {
		case "INCREMENT":
			return store.GetState("count").(int) + 1
		default:
			return store.GetAllState()
		}
	}
	store.Reductor(reductor)
	value := store.Dispatch("INCREMENT")
	if value != 2 {
		t.Error("The value are different. Action not correct.")
	}
}
