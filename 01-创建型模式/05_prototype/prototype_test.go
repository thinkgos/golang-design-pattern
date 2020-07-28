package prototype

import "testing"

func TestCloneFromManager(t *testing.T) {
	manager := NewPrototypeManager()
	manager.Set("t1", &Type1{name: "type1"})

	t1 := manager.Get("t1").Clone()
	if t1.Name() != "type1" {
		t.Fatal("error")
	}
}
