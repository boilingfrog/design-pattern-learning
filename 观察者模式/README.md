<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [观察者模式](#%E8%A7%82%E5%AF%9F%E8%80%85%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 观察者模式

### 定义

观察者模式(Observer Design Pattern)定义了一种一对多的依赖关系，让多个观察者对象同时监听一个主题对象。这个主题对象在状态发生变化的时，会通知所有的观察者对象，使他们能够更新自己。  

定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。  

### 适用场景

1、当一个对象状态的改变需要改变其他对象，或实际对象是事先未知的或动态变化的时，可使用观察者模式；  

2、当应用中的一些对象必须观察其他对象时，可使用该模式。但仅能在有限时间内或特定情况下使用。

### 优点

1、降低了目标与观察者之间的耦合关系，两者之间是抽象耦合关系。  

2、目标与观察者之间建立了一套触发机制。  

### 缺点

1、目标与观察者之间的依赖关系并没有完全解除，而且有可能出现循环引用。    

2、当观察者对象很多时，通知的发布会花费很多时间，影响程序的效率。  

### 代码实现

被观察者有信息更新的时候，通知到所有的观察者。   

```go
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type Customer struct {
	name string
}

func NewCustomer(name string) *Customer {
	return &Customer{
		name: name,
	}
}

func (r *Customer) Update(s *Subject) {
	fmt.Printf("%s received %s\n", r.name, s.context)
}
```

测试代码  

```go
func TestObserver(t *testing.T) {
	subject := NewSubject()
	reader1 := NewCustomer("小明")
	reader2 := NewCustomer("小红")
	reader3 := NewCustomer("小李")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	for i := 1; i <= 10; i++ {
		subject.UpdateContext(fmt.Sprintf("更新了%d", i))
		fmt.Println("+++++++++++++++++++++++++++++++++")
	}
}
```

结构图  

<img src="/img/pattern-observer.png" alt="observer" />

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/观察者模式   
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001    
【golang-design-pattern】https://github.com/senghoo/golang-design-pattern    


