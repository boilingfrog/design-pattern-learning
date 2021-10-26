package main

import (
	"fmt"
)

func main() {
	payment := NewPayment("小明", 12, &Weixin{})
	payment.Pay()
	fmt.Println()
	apay := NewPayment("小明", 12, &Ali{})
	apay.Pay()
}

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Account string
	Money   int
}

func NewPayment(account string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Account: account,
			Money:   money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type Weixin struct{}

func (*Weixin) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay %d元 to %s by weixin", ctx.Money, ctx.Account)
}

type Ali struct{}

func (*Ali) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay %d元 to %s by ali", ctx.Money, ctx.Account)
}
