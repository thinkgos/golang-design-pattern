package state

func ExampleWeek() {
	week := []Week{new(Sunday), new(Monday), new(Tuesday), new(Wednesday), new(Thursday), new(Friday), new(Saturday)}
	ctx := NewDayContext()

	for _, v := range week {
		ctx.SetDay(v) // 改变状态
		ctx.Today()   // 查看行为
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
