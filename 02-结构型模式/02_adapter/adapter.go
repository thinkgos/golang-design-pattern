package adapter

// Target 是适配的目标接口,暴露给用户最终接口
type Target interface {
	Request() string
}

// Provider 是被适配的目标接口
type Provider interface {
	SpecificRequest() string
}

// providerImpl 是被适配的目标类,提供给adapter的被适配接口
type providerImpl struct{}

// NewProvider 是被适配接口的工厂函数
func NewProvider() Provider {
	return &providerImpl{}
}

//SpecificRequest 是目标类的一个方法
func (*providerImpl) SpecificRequest() string {
	return "Provider method"
}

//adapter 是转换Provider为Target接口的适配器
type adapter struct {
	Provider
}

//NewAdapter 是adapter的工厂函数
func NewAdapter(p Provider) Target {
	return &adapter{p}
}

// Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
