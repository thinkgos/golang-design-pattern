package facade

import (
	"fmt"
)

func Draw(s Shape) {
	fmt.Println(s.Draw())
}

func ExampleFacade() {
	rect := NewRectangle()
	Draw(rect)
	circle := NewCircle()
	Draw(circle)
	// output:
	// Rectangle Draw
	// Circle Draw
}
