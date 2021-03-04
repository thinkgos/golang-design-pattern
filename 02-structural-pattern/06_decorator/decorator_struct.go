package decorator

type Component interface {
	Calc() int
}

type BaseComponent struct{}

func (*BaseComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{c, num}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{c, num}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}
