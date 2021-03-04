package facade

// Shape is facade interface of facade package
// 提供给客户端的统一接口
type Shape interface {
	Draw() string
}

type Rectangle struct{}

// NewRectangle ...
func NewRectangle() Shape {
	return &Rectangle{}
}

func (*Rectangle) Draw() string {
	return "Rectangle Draw"
}

type Circle struct{}

// NewCircle ...
func NewCircle() Shape {
	return &Circle{}
}

func (*Circle) Draw() string {
	return "Circle Draw"
}
