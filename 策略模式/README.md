<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [策略模式](#%E7%AD%96%E7%95%A5%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [使用场景](#%E4%BD%BF%E7%94%A8%E5%9C%BA%E6%99%AF)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 策略模式

### 定义

策略模式定义了算法家族，分别封装起来，让他们之间可以相互替换，此模式让算法的变化，不会影响到客户端的使用，也称为政策模式(Policy)。  

策略模式主要的作用还是解耦策略的定义、创建和使用，控制代码的复杂度，让每个部分都不至于过于复杂、代码量过多。除此之外，对于复杂代码来说，策略模式还能让其满足开闭原则，添加新策略的时候，最小化、集中化代码改动，减少引入bug的风险。  

### 优点

1、算法可以自由切换。   

2、避免使用多重条件判断。 
  
3、扩展性良好。  

### 缺点

1、策略类会增多。  

2、所有策略类都需要对外暴露。  

### 使用场景

1、如果在一个系统里面有许多类，它们之间的区别仅在于它们的行为，那么使用策略模式可以动态地让一个对象在许多行为中选择一种行为。   

2、一个系统需要动态地在几种算法中选择一种。   

3、如果一个对象有很多的行为，如果不用恰当的模式，这些行为就只好使用多重的条件选择语句来实现。  

### 代码实现

比如我们在做支付的项目中，对于支付的类型，我们可能会有多种类型，我们可能要做微信支付，支付宝支付，苹果内购等等一些列的产品。。。  

所以我们可以把每种支付产品定义成一种策略，然后根据不同的业务场景选择不同的支付产品     

```go
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
```
