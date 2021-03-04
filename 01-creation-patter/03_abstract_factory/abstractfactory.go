package abstractfactory

import "fmt"

// AbstractProductOrderMain 为订单主记录
type AbstractProductOrderMain interface {
	SaveOrderMain()
}

// AbstractProductOrderDetail 为订单详情纪录
type AbstractProductOrderDetail interface {
	SaveOrderDetail()
}

// AbstractFactory 抽象模式工厂接口
type AbstractFactory interface {
	CreateOrderMain() AbstractProductOrderMain
	CreateOrderDetail() AbstractProductOrderDetail
}

/*******************************例1***********************************************/

// RDBMainDAP 具体产品类 为关系型数据库的OrderMainDAO实现
type RDBMain struct{}

// SaveOrderMain ...
func (*RDBMain) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// RDBDetail 具体产品类 为关系型数据库的OrderDetailDAO实现
type RDBDetail struct{}

// SaveOrderDetail ...
func (*RDBDetail) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBFactory 具体工厂类 RDB抽象工厂实现
type RDBFactory struct{}

func (*RDBFactory) CreateOrderMain() AbstractProductOrderMain {
	return new(RDBMain)
}

func (*RDBFactory) CreateOrderDetail() AbstractProductOrderDetail {
	return new(RDBDetail)
}

/*********************************例2*********************************************/

// XMLMain 具体产品类 XML存储
type XMLMain struct{}

// SaveOrderMain ...
func (*XMLMain) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// XMLDetail 具体产品类 XML存储
type XMLDetail struct{}

// SaveOrderDetail ...
func (*XMLDetail) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

// XMLFactory 具体工厂类 RDB抽象工厂实现
type XMLFactory struct{}

func (*XMLFactory) CreateOrderMain() AbstractProductOrderMain {
	return new(XMLMain)
}

func (*XMLFactory) CreateOrderDetail() AbstractProductOrderDetail {
	return new(XMLDetail)
}
