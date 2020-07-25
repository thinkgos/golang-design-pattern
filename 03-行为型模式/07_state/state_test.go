package state

func ExampleWeek() {
	week := []Week{new(Sunday), new(Monday), new(Tuesday), new(Wednesday), new(Thursday), new(Friday), new(Saturday)}
	ctx := NewDayContext()

	for _, v := range week {
		ctx.SetDay(v)
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
