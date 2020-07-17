package visitor

func ExampleVisitor() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Accept(&ServiceRequestVisitor{})
	c.Accept(&AnalysisVisitor{})
	// Output:
	// serving enterprise customer A company
	// serving enterprise customer B company
	// analysis enterprise customer A company
	// analysis enterprise customer B company
}
