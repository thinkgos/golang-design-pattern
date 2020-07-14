package decorator

import (
	"fmt"
)

func ExampleDecoratorFunc() {
	var c = NewBase()
	c = WarpAddDecoratorFunc(c, 10)
	c = WarpMulDecoratorFunc(c, 8)
	res := c()

	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}
