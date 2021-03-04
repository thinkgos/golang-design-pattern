package interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	p := new(Parser)

	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()

	if expect := 1; res != expect {
		t.Fatalf("expect %d got %d", expect, res)
	}
}
