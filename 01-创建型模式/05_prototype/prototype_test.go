package prototype

import "testing"

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

func (t *Type1) Name() string {
	return t.name
}

func TestCloneFromManager(t *testing.T) {
	manager := NewPrototypeManager()
	manager.Set("t1", &Type1{name: "type1"})

	t1 := manager.Get("t1").Clone()
	if t1.Name() != "type1" {
		t.Fatal("error")
	}
}
