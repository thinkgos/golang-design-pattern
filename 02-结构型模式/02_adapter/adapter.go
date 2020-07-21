package adapter

// Target 是适配的目标接口,暴露给用户最终接口
type Target interface {
	Request() string
}

// Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// adapteeImpl
// 是被适配的目标类,提供给adapter的被适配接口
type adapteeImpl struct{}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "Adaptee method"
}

// adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

// NewAdapter 是adapter的工厂函数
func NewAdapter(p Adaptee) Target {
	return &adapter{p}
}

// Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
