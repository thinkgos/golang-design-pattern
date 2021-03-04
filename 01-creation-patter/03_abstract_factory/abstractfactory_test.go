package abstractfactory

func getMainAndDetail(factory AbstractFactory) {
	factory.CreateOrderMain().SaveOrderMain()
	factory.CreateOrderDetail().SaveOrderDetail()
}

func ExampleRdbFactory() {
	var factory AbstractFactory
	factory = &RDBFactory{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

func ExampleXmlFactory() {
	var factory AbstractFactory
	factory = &XMLFactory{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}
