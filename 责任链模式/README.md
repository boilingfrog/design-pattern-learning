<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [责任链模式](#%E8%B4%A3%E4%BB%BB%E9%93%BE%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [责任链模式对比装饰模式](#%E8%B4%A3%E4%BB%BB%E9%93%BE%E6%A8%A1%E5%BC%8F%E5%AF%B9%E6%AF%94%E8%A3%85%E9%A5%B0%E6%A8%A1%E5%BC%8F)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 责任链模式

### 定义

责任链模式(Chain Of Responsibility):使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的藕合关系。将这个对象连成一条链，并沿着这条链传递该请求，直到一个对象处理完为止。   

是一种处理请求的模式，它让多个处理器都有机会处理该请求，直到其中某个处理成功为止。责任链模式把多个处理器串成链，然后让请求在链上传递。   

举个栗子  

比如学生请假，小于等于 2 天班主任可以审批；小于等于 7 天，系主任可以审批；小于或等于 10 天，院长可以批准。其他情况，不予审批。   

用责任链模式设计审批流程，每个审核者只关心自己责任范围内的请求，并且处理它。对于超出自己责任范围的，扔给下一个审核者处理，这样，将来继续添加审核者的时候，不用改动现有逻辑。  

<img src="/img/pattern-processor.png" alt="responsibility" />  

### 优点

1、降低耦合度。它将请求的发送者和接收者解耦；  

2、简化了对象。使得对象不需要知道链的结构；  

3、增强给对象指派职责的灵活性。通过改变链内的成员或者调动它们的次序，允许动态地新增或者删除责任；   

4、增加新的请求处理类很方便。   

### 缺点

不能保证请求一定被接收   

### 适用范围

1、多个对象可以处理一个请求，但具体由哪个对象处理该请求在运行时自动确定；  

2、可动态指定一组对象处理请求，或添加新的处理者；  

3、需要在不明确指定请求处理者的情况下，向多个处理者中的一个提交请求。  

### 代码实现

上面学生请假的栗子：  

学生请假，小于等于 2 天班主任可以审批；小于等于 7 天，系主任可以审批；小于或等于 10 天，院长可以批准。其他情况，不予审批。   

```go
type Teacher interface {
	HaveRight(day int) bool
	HandleApproveRequest(name string, day int) bool
}

type RequestChain struct {
	Teacher
	approver *RequestChain
}

func (r *RequestChain) SetApprover(m *RequestChain) {
	r.approver = m
}

func (r *RequestChain) HandleApproveRequest(name string, day int) bool {
	if r.Teacher.HaveRight(day) {
		return r.Teacher.HandleApproveRequest(name, day)
	}
	if r.approver != nil {
		return r.approver.HandleApproveRequest(name, day)
	}
	fmt.Println("请假时间太久了，不予审批")
	return false
}

func (r *RequestChain) HaveRight(day int) bool {
	return true
}

type HeadTeacher struct{}

func NewHeadTeacherChain() *RequestChain {
	return &RequestChain{
		Teacher: &HeadTeacher{},
	}
}

func (*HeadTeacher) HaveRight(day int) bool {
	return day <= 2
}

func (*HeadTeacher) HandleApproveRequest(name string, day int) bool {
	fmt.Println(fmt.Sprintf("班主任，批准了%s的请假申请,请假天数%d", name, day))
	return true
}

type DepTeacher struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Teacher: &DepTeacher{},
	}
}

func (*DepTeacher) HaveRight(day int) bool {
	return day <= 7
}

func (*DepTeacher) HandleApproveRequest(name string, day int) bool {
	fmt.Println(fmt.Sprintf("系主任，批准了%s的请假申请,请假天数%d", name, day))
	return true
}

type DeanTeacher struct{}

func NewDeanTeacherChain() *RequestChain {
	return &RequestChain{
		Teacher: &DeanTeacher{},
	}
}

func (*DeanTeacher) HaveRight(day int) bool {
	return day <= 10
}

func (*DeanTeacher) HandleApproveRequest(name string, day int) bool {
	fmt.Println(fmt.Sprintf("院长，批准了%s的请假申请,请假天数%d", name, day))
	return true
}
```

测试代码  

```go
func TestApproveChain(t *testing.T) {
	c1 := NewHeadTeacherChain()
	c2 := NewDepManagerChain()
	c3 := NewDeanTeacherChain()

	c1.SetApprover(c2)
	c2.SetApprover(c3)

	var c Teacher = c1
	assert.Equal(t, true, c.HandleApproveRequest("小明", 3))
	assert.Equal(t, true, c.HandleApproveRequest("小红", 2))
	assert.Equal(t, false, c.HandleApproveRequest("小龙", 30))
}
```

输出  

```
系主任，批准了小明的请假申请,请假天数3
班主任，批准了小红的请假申请,请假天数2
请假时间太久了，不予审批
```

### 责任链模式对比装饰模式

责任链和装饰模式非常相似。两者都依赖递归组合将需要执行的操作传递给一系列对象。来看下两者的不同点。   

1、责任链的管理者可以相互独立地执行一切操作，还可以随时停止传递请求。  

2、装饰模式可以对对象进行扩展，但是不能中断请求的传递。   

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/责任链模式  
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   
【责任链模式（责任链模式）详解】http://c.biancheng.net/view/1383.html   
【责任链模式】https://boilingfrog.github.io/2021/11/22/%E4%BD%BF%E7%94%A8go%E5%AE%9E%E7%8E%B0%E8%B4%A3%E4%BB%BB%E9%93%BE%E6%A8%A1%E5%BC%8F/  