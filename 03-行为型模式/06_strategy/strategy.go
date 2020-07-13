package strategy

import "fmt"

// 策略接口
type Strategy interface {
	Pay(*PaymentContext)
}

// 策略上下文
type PaymentContext struct {
	Name   string
	CardID string
	Money  int
}

type Payment struct {
	context  *PaymentContext
	strategy Strategy
}

// 根据不同的策创建实例方法
func NewPayment(name, cardId string, money int, strategy Strategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

// 调策略
func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

// 策略实例
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

// 策略实例
type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}
