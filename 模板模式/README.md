<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [模板模式](#%E6%A8%A1%E6%9D%BF%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [模板模式的作用](#%E6%A8%A1%E6%9D%BF%E6%A8%A1%E5%BC%8F%E7%9A%84%E4%BD%9C%E7%94%A8)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [回调](#%E5%9B%9E%E8%B0%83)
    - [模板模式 VS 回调](#%E6%A8%A1%E6%9D%BF%E6%A8%A1%E5%BC%8F-vs-%E5%9B%9E%E8%B0%83)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 模板模式

### 定义

模板模式(TemplateMethod):定义一个操作中的算法骨架，而将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重新定义该算法的某些特定步骤。

模板方法模式就是提供一个代码复用平台，当不变和可变的行为在方法的子类实现中混合在一起的时候，不变的行为就会在子类中重复出现。通过模板方法模式把这些行为搬到单一的地方，这样就帮助子类摆脱重复的不变行为的纠缠。  

### 模板模式的作用  

1、复用

所有的子类可以复用父类中提供的模板方法的代码  

2、扩展

框架通过模板模式提供功能扩展点，让框架用户可以在不修改框架源码的情况下，基于扩展点定制化框架的功能   
   
### 优点

1、封装不变部分，扩展可变部分；   

2、提取公共代码，便于维护；   

3、行为由父类控制，子类实现。

### 缺点

每一个不同的实现都需要一个子类来实现，导致类的个数增加，使得系统更加庞大。

### 适用范围

1、有多个子类共有的方法，且逻辑相同；  

2、重要的、复杂的方法，可以考虑作为模板方法。  

### 代码实现

考试时候试卷，对于试题部分，同一场考试内容都是一样的。试卷做完交卷只是我们每人个人填写的答案不同，那么试题就可以作为模板，我们只用去写答案。  

```go
type TestPaperImpl interface {
	testQuestion1()
	testQuestion2()
	Answer1()
	Answer2()
}

type testPaper struct {
}

func (t *testPaper) testQuestion1() {
	fmt.Println("问题：中国有多少个民族")
}

func (t *testPaper) testQuestion2() {
	fmt.Println("问题：中国有多大")
}

func (t *testPaper) Answer1() {
}

func (t *testPaper) Answer2() {
}

// 封装具体步骤
func doPaper(paper TestPaperImpl) {
	paper.testQuestion1()
	paper.Answer1()

	paper.testQuestion2()
	paper.Answer2()
}

type student1 struct {
	*testPaper
}

func (s *student1) Answer1() {
	fmt.Println("答案：56")
}

func (s *student1) Answer2() {
	fmt.Println("答案：很大")
}

type student2 struct {
	*testPaper
}

func (s *student2) Answer1() {
	fmt.Println("答案：51")
}

func (s *student2) Answer2() {
	fmt.Println("答案：不知道")
}
```

测试文件  

```go
func TestTestPaper(t *testing.T) {
	st1 := &student1{}
	doPaper(st1)

	fmt.Println("++++++++++++++")
	st2 := &student2{}
	doPaper(st2)
}
```

结果  

```
问题：中国有多少个民族
答案：56
问题：中国有多大
答案：很大
++++++++++++++
问题：中国有多少个民族
答案：51
问题：中国有多大
答案：不知道
```

结构图  

<img src="/img/pattern-template.png" alt="template" />

### 回调  

回调起到的作用和模板模式一样  

相对于普通的函数调用来说，回调是一种双向调用关系。A类事先注册某个函数F到B类，A类在调用B类的P函数的时候，B类反过来调用A类注册给它的F函数。这里的F函数就是“回调函数”。A调用B，B反过来又调用A，这种调用机制就叫作“回调”  

回调分为两种：  

1、同步回调  

在函数返回之前执行回调函数，同步回调看起来有点像模板模式  

2、异步回调

在函数返回之后执行回调函数  

如果做过支付的同学肯定很熟悉这个，例如微信支付，我们调用微信支付进行付款，成功之后我们的服务器会收到微信端支付的消息回调，然后进行支付成功之后的后续操作。  

上面的考试例子使用回调实现  

```go
type testPaperCallback struct {
}

func (t *testPaperCallback) testQuestion1() {
	fmt.Println("问题1：中国有多少个民族")
}

func (t *testPaperCallback) testQuestion2() {
	fmt.Println("问题2：中国有多大")
}

func (t *testPaperCallback) Answer(callback CallbackImpl) {
	t.testQuestion1()
	t.testQuestion2()
	callback.AnswerCallback()
}

type CallbackImpl interface {
	AnswerCallback()
}

type student3 struct {
	*testPaperCallback
}

func (s *student3) AnswerCallback() {
	fmt.Println("答案1：56")
	fmt.Println("答案2：测试")
}

func doPaperCallback(student *student3) {
	student.Answer(&student3{})
}
```

结构图  

<img src="/img/pattern-callback.png" alt="callback" />

#### 模板模式 VS 回调

从应用场景上来看，同步回调跟模板模式几乎一致。它们都是在一个大的算法骨架中，自由替换其中的某个步骤，起到代码复用和扩展的目的。而异步回调跟模板模式有较大差别，更像是观察者模式。  

从代码实现上来看，回调和模板模式完全不同。回调基于组合关系来实现，把一个对象传递给另一个对象，是一种对象之间的关系；模板模式基于继承关系来实现，子类重写父类的抽象方法，是一种类之间的关系。  

- 回调可以使用匿名类来创建回调对象，可以不用事先定义类；而模板模式针对不同的实现都要定义不同的子类。  

- 如果某个类中定义了多个模板方法，每个方法都有对应的抽象方法，那即便我们只用到其中的一个模板方法，子类也必须实现所有的抽象方法。而回调就更加灵活，我们只需要往用到的模板方法中注入回调对象即可。  

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/模板模式   
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001    
【模板模式】https://boilingfrog.github.io/2021/11/20/%E4%BD%BF%E7%94%A8go%E5%AE%9E%E7%8E%B0%E6%A8%A1%E6%9D%BF%E6%A8%A1%E5%BC%8F/  