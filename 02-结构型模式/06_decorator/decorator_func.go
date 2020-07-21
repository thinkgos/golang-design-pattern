package decorator

type ComponentFunc func() int

func NewEmptyBase() ComponentFunc {
	return func() int {
		return 0
	}
}

func WarpMulDecoratorFunc(c ComponentFunc, num int) ComponentFunc {
	return func() int {
		return c() * num
	}
}

func WarpAddDecoratorFunc(c ComponentFunc, num int) ComponentFunc {
	return func() int {
		return c() + num
	}
}
