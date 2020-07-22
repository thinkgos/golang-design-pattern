package visitor

func ExampleVisitor() {
	ac := NewEnterpriseCustomer("A company")
	bc := NewEnterpriseCustomer("B company")
	vis1 := &ServiceRequestVisitor{}
	vis2 := &AnalysisVisitor{}
	ac.Accept(vis1)
	bc.Accept(vis1)
	ac.Accept(vis2)
	bc.Accept(vis2)
	// Output:
	// serving enterprise customer A company
	// serving enterprise customer B company
	// analysis enterprise customer A company
	// analysis enterprise customer B company
}
