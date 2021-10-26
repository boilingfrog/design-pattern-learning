package main

import (
	"fmt"
)

func main() {
	payment := NewPayment(&Weixin{
		&PaymentConf{
			appId:     "wx1323234343434",
			notifyURL: "weixin.notifyURL.com",
		},
	})
	payment.Pay("小明", 12)
	fmt.Println()
	apay := NewPayment(&Ali{
		&PaymentConf{
			appId:     "al1323234343434",
			notifyURL: "ali.notifyURL.com",
		},
	})
	apay.Pay("小红", 16)
}

type Context struct {
	strategy PaymentStrategy
}

type PaymentConf struct {
	appId     string
	notifyURL string
}

func NewPayment(strategy PaymentStrategy) *Context {
	return &Context{
		strategy: strategy,
	}
}

func (p *Context) Pay(account string, money int) {
	p.strategy.Pay(account, money)
}

type PaymentStrategy interface {
	Pay(account string, money int)
}

type Weixin struct {
	*PaymentConf
}

func (w *Weixin) Pay(account string, money int) {
	fmt.Printf("Pay %d元 to %s by weixin", money, account)
}

type Ali struct {
	*PaymentConf
}

func (a *Ali) Pay(account string, money int) {
	fmt.Printf("Pay %d元 to %s by ali", money, account)
}
