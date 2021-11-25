## 访问者模式

### 定义

访问者模式(Visitor):表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变元素类的前提下定义作用于这些元素的新操作。  

### 优点

### 缺点

### 适用范围

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

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/访问者模式    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   