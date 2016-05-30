package godux

import "testing"

func setStoreState(t *testing.T) {
	store := NewStore()
	store.Setstate("count", 1)
	if store.GetState("count") != 1 {
		t.Error("Store state are different")
	}
}
