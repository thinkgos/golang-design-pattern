package proxy

// 代理和被代理之间的抽象接口
type Subject interface {
	Do() string
}

// RealSubject 被代理对像
type RealSubject struct{}

// Do 实现抽像接口
func (RealSubject) Do() string {
	return "real"
}

// Proxy 代理对象
type Proxy struct {
	real Subject
}

// NewProxy 创建代理
func NewProxy(rs Subject) *Proxy {
	return &Proxy{rs}
}

// Do 实现抽象接口
func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}
