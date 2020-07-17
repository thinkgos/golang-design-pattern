package visitor

import "fmt"

// 被访问者实现的暴露的接口.
type Customer interface {
	Accept(Visitor)
	Name() string
}

// 访问者实现的方法,也就是访问方法
type Visitor interface {
	Visit(Customer)
}

type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

// 被访问者1
type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

func (c EnterpriseCustomer) Name() string {
	return c.name
}

// 访问者1
type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	fmt.Printf("serving enterprise customer %s\n", customer.Name())
}

// 访问者2
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	fmt.Printf("analysis enterprise customer %s\n", customer.Name())
}
