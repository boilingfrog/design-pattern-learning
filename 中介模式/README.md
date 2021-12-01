<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [中介模式](#%E4%B8%AD%E4%BB%8B%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 中介模式

### 定义

中介模式(Mediator):用一个中介对象来封装一系列的对象交互。中介者使个各对象不需要显示的相互引用，从而使其藕合松散，而且可以独立的改变它们之间的交互。   

中介模式的设计思想跟中间层很像，通过引入中介这个中间层，将一组对象之间的交互关系（或者说依赖关系）从多对多（网状关系）转换为一对多（星状关系）。原来一个对象要跟n个对象交互，现在只需要跟一个中介对象交互，从而最小化对象之间的交互关系，降低了代码的复杂度，提高了代码的可读性和可维护性。    

### 优点

1、可以减轻应用中多个组件间的耦合情况；   

2、降低了组件的复杂度，将一对多转化成了一对一；   

### 缺点

中介类有可能会变成大而复杂的“上帝类”（God Class）。  

### 适用范围

系统中对象之间存在比较复杂的引用关系，导致它们之间的依赖关系结构混乱而且难以复用该对象。  

### 代码实现

```go
type Country interface {
	SendMess(message string)
	GetMess(message string)
}

type USA struct {
	mediator *UnitedNationsSecurityCouncil
}

func (usa *USA) SendMess(message string) {
	usa.mediator.ForwardMessage(usa, message)
}

func (usa *USA) GetMess(message string) {
	fmt.Println("USA 获得对方的消息：", message)
}

type Irap struct {
	mediator *UnitedNationsSecurityCouncil
}

func (ir *Irap) SendMess(message string) {
	ir.mediator.ForwardMessage(ir, message)
}

func (ir *Irap) GetMess(message string) {
	fmt.Println("Irap 获得对方的消息：", message)
}

type Mediator interface {
	ForwardMessage(country Country, message string)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Irap
}

func (uns *UnitedNationsSecurityCouncil) ForwardMessage(country Country, message string) {
	switch country.(type) {
	case *USA:
		uns.Irap.GetMess(message)
	case *Irap:
		uns.USA.GetMess(message)
	default:
		fmt.Println("国家不在联合国")
	}
}
```

测试代码  

```go
func TestMediator_ForwardMessage(t *testing.T) {
	// 创建中介者，联合国
	mediator := &UnitedNationsSecurityCouncil{}

	usa := USA{mediator}
	mediator.USA = usa
	usa.SendMess("不准研制核武器，否则要发动战争了")

	irap := Irap{mediator}
	mediator.Irap = irap
	irap.SendMess("我们没有核武器")
}
```

输出结果  

```
Irap 获得对方的消息： 不准研制核武器，否则要发动战争了
USA 获得对方的消息： 我们没有核武器
```

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/中介模式    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   