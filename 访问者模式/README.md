<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [访问者模式](#%E8%AE%BF%E9%97%AE%E8%80%85%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [什么是 Double Dispatch](#%E4%BB%80%E4%B9%88%E6%98%AF-double-dispatch)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 访问者模式

### 定义

访问者模式(Visitor):表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变元素类的前提下定义作用于这些元素的新操作。  

使用访问者模式，元素的执行算法可以随着访问者改变而改变。主要意图是将数据结构与数据操作分离。   

不过作为比较难理解的设计模式之一，因为它难理解、难实现，应用它会导致代码的可读性、可维护性变差，所以，访问者模式在实际的软件开发中很少被用到，在没有特别必要的情况下，访问者模式是不建议使用的。     

### 优点

1、开闭原则。 你可以引入在不同类对象上执行的新行为， 且无需对这些类做出修改。  

2、单一职责原则。 可将同一行为的不同版本移到同一个类中。  

3、灵活性更好。   

### 缺点

1、具体元素变更比较困难。每次在元素层次结构中添加或移除一个类时，都要更新所有的访问者。  

2、比较难理解，应用它会导致代码的可读性、可维护性变差。  

### 适用范围

1、对象结构中对象对应的类很少改变，但经常需要在此对象结构上定义新的操作。   

2、需要对一个对象结构中的对象进行很多不同的并且不相关的操作，而需要避免让这些操作"污染"这些对象的类，也不希望在增加新操作时修改这些类。  

### 代码实现

代码实现：  

```go
type Visitor interface {
	VisitConcreteElementA(cea *ConcreteElementA)
	VisitConcreteElementB(ceb *ConcreteElementB)
}

type ConcreteVisitor1 struct {
}

func (cea *ConcreteVisitor1) VisitConcreteElementA(concreteElementA *ConcreteElementA) {
	fmt.Println("concreteVisitor1 visitConcreteElementA")
}

func (*ConcreteVisitor1) VisitConcreteElementB(concreteElementB *ConcreteElementB) {
	fmt.Println("concreteVisitor1 visitConcreteElementB")
}

type ConcreteVisitor2 struct {
}

func (*ConcreteVisitor2) VisitConcreteElementA(concreteElementA *ConcreteElementA) {
	fmt.Println("concreteVisitor2 visitConcreteElementA")
}

func (*ConcreteVisitor2) VisitConcreteElementB(concreteElementB *ConcreteElementB) {
	fmt.Println("concreteVisitor2 visitConcreteElementB")
}

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct {
}

func (cea *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(cea)
}

type ConcreteElementB struct {
}

func (ceb *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ceb)
}
```

测试代码：  

```go
func TestVisitor(t *testing.T) {
	var elements []Element
	elements = append(elements, &ConcreteElementA{})
	elements = append(elements, &ConcreteElementB{})

	for _, item := range elements {
		cv1 := &ConcreteVisitor1{}
		cv2 := &ConcreteVisitor2{}
		item.Accept(cv1)
		item.Accept(cv2)
	}
}
```

结构图：   

<img src="/img/pattern-visitor.png" alt="visitor" />  

### 什么是 Double Dispatch  

什么是分派？  

分派即 Dispatch，在面向对象编程语言中，我们可以把方法调用理解为一种消息传递（Dispatch）。一个对象调用另一个对象的方法，相当于给被调用对象发送一个消息，这个消息包括对象名、方法名、方法参数等信息。  

什么是单分派？  

单分派，即执行哪个对象的方法，根据对象的运行时类型决定；执行对象的哪个方法，根据方法参数的编译时类型决定。  

什么是双分派？  

双分派，即执行哪个对象的方法，根据对象的运行时类型来决定；执行对象的哪个方法，根据方法参数的运行时的类型来决定。  

具体到编程语言的语法机制，Single Dispatch 和 Double Dispatch 跟多态和函数重载直接相关。  

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/访问者模式    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   
【双分派-访问者模式的前世今生】https://www.codenong.com/cs110749395/    