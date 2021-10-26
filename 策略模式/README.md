<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [策略模式](#%E7%AD%96%E7%95%A5%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [使用场景](#%E4%BD%BF%E7%94%A8%E5%9C%BA%E6%99%AF)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [策略模式和工厂模式的区别](#%E7%AD%96%E7%95%A5%E6%A8%A1%E5%BC%8F%E5%92%8C%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F%E7%9A%84%E5%8C%BA%E5%88%AB)
  - [参考](#%E5%8F%82%E8%80%83)

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
```

来看下结构图  

<img src="/img/uml-strategy.png" alt="strategy" />

### 策略模式和工厂模式的区别

**工厂模式**

1、目的是创建不同且相关的对象  
 
2、侧重于"创建对象"  

3、实现方式上可以通过父类或者接口  

4、一般创建对象应该是现实世界中某种事物的映射，有它自己的属性与方法  

**策略模式**

1、目的实现方便地替换不同的算法类  

2、侧重于算法(行为)实现  

3、实现主要通过接口  

4、创建对象对行为的抽象而非对对象的抽象，很可能没有属于自己的属性  

### 参考

【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】设计模式之美    
【工厂模式】https://www.cnblogs.com/ricklz/p/15399178.html      




