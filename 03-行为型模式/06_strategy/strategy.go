package strategy

import "fmt"

type Strategy interface {
	Pay(*PaymentContext)
}

type Payment struct {
	context  *PaymentContext
	strategy Strategy
}

type PaymentContext struct {
	Name   string
	CardID string
	Money  int
}

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

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}
