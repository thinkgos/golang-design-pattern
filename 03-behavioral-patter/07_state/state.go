package state

import "fmt"

type tWeek byte

// 状态接口
type Week interface {
	Today()
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

// 行为
func (d *DayContext) Today() {
	d.today.Today()
}

// 改变状态
func (d *DayContext) SetDay(w Week) {
	d.today = w
}

type Sunday struct{}

func (*Sunday) Today() {
	fmt.Printf("Sunday\n")
}

type Monday struct{}

func (*Monday) Today() {
	fmt.Printf("Monday\n")
}

type Tuesday struct{}

func (*Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

type Wednesday struct{}

func (*Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

type Thursday struct{}

func (*Thursday) Today() {
	fmt.Printf("Thursday\n")
}

type Friday struct{}

func (*Friday) Today() {
	fmt.Printf("Friday\n")
}

type Saturday struct{}

func (*Saturday) Today() {
	fmt.Printf("Saturday\n")
}
