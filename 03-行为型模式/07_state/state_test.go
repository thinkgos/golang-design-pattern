package state

func ExampleWeek() {
	ctx := NewDayContext()

	for i := 0; i < 7; i++ {
		ctx.SetDay(tWeek(i))
		ctx.Today()
	}
	// Output:
	// Sunday
	// Monday
	// Tuesday
	// Wednesday
	// Thursday
	// Friday
	// Saturday
}
