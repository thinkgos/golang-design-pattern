package strategy

func ExamplePayByCash() {
	payment := NewContext(&Cash{})
	payment.Do()
	// Output:
	// Do by cash
}

func ExamplePayByBank() {
	payment := NewContext(&Bank{})
	payment.Do()
	// Output:
	// Do by bank account
}
