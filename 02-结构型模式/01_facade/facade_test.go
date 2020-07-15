package facade

import "testing"

// TestFacadeAPI ...
func TestFacadeAPI(t *testing.T) {
	expect := "A module running\nB module running"
	api := NewAPI()
	ret := api.Test()

	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}
