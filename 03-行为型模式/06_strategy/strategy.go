package strategy

import "fmt"

// 抽象的策略
type Strategy interface {
	Do()
}

// 实现一个上下文的类
type Context struct {
	Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{
		strategy,
	}
}

// 策略1
type Cash struct{}

func (*Cash) Do() {
	fmt.Printf("Do by cash")
}

// 策略2
type Bank struct{}

func (*Bank) Do() {
	fmt.Printf("Do by bank account")

}
