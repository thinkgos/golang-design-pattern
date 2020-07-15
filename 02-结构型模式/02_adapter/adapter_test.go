package adapter

import "testing"

func TestAdapter(t *testing.T) {
	expect := "Provider method"
	provider := NewProvider()
	target := NewAdapter(provider)

	if res := target.Request(); res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
