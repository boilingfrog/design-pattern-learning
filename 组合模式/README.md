<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [组合模式](#%E7%BB%84%E5%90%88%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 组合模式

### 定义

组合模式(Composite),将对象组合成树形结构以表示'部分-整体'的层次关系。组合模式使得用户对单个对象和组合对象的使用具有一致性。    

组合模式，将一组对象组织成树形结构，将单个对象和组合对象都看做树中的节点，以统一处理逻辑，并且它利用树形结构的特点，递归地处理每个子树，依次简化代码实现。使用组合模式的前提在于，你的业务场景必须能够表示成树形结构。所以，组合模式的应用场景也比较局限，它并不是一种很常用的设计模式。  

### 适用范围

 1、您想表示对象的部分-整体层次结构（树形结构）；  
 
 2、您希望用户忽略组合对象与单个对象的不同，用户将统一地使用组合结构中的所有对象。  

### 优点

1、高层模块调用简单；  
 
2、节点自由增加。  

### 缺点

1、设计较复杂，客户端需要花更多时间理清类之间的层次关系；  

2、不容易限制容器中的构件；  

3、不容易用继承的方法来增加构件的新功能；  

### 代码实现

公司有部门，部门下面有员工，使用组合模式来表示这种层次关系。   

```go
type Employee struct {
	no           int
	name         string
	dept         string
	salary       int
	subordinates []*Employee
}

func NewEmployee(no int, name, dept string, salary int) *Employee {
	return &Employee{
		no:           no,
		name:         name,
		dept:         dept,
		salary:       salary,
		subordinates: []*Employee{},
	}
}

func (e *Employee) add(em *Employee) {
	e.subordinates = append(e.subordinates, em)
}

func (e *Employee) remove(em *Employee) {
	for index, item := range e.subordinates {
		if item.no == em.no {
			e.subordinates = e.subordinates[:index+copy(e.subordinates[index:], e.subordinates[index+1:])]
			break
		}
	}
}

func (e *Employee) getSubordinates() []*Employee {
	return e.subordinates
}

func (e *Employee) toString() {
	fmt.Println(fmt.Sprintf("Employee:no:%d;name:%s;dept:%s;salary:%d。", e.no, e.name, e.dept, e.salary))
}
```

测试文件  

```go
func TestComposite(t *testing.T) {
	ceo := NewEmployee(1, "马云", "CEO", 1000000)

	headMarketing := NewEmployee(2, "小白", "市场总监", 50000)
	clerk1 := NewEmployee(1001, "小明", "Marketing", 10000)
	clerk2 := NewEmployee(1002, "小张", "Marketing", 10000)

	cto := NewEmployee(3, "小龙", "CTO", 90000)
	tc1 := NewEmployee(1003, "马龙", "Technology", 10000)
	tc2 := NewEmployee(1004, "张龙", "Technology", 10000)

	ceo.add(headMarketing)
	headMarketing.add(clerk1)
	headMarketing.add(clerk2)

	ceo.add(cto)
	cto.add(tc1)
	cto.add(tc2)

	ceo.toString()
	for _, subordinate := range ceo.subordinates {
		fmt.Println("部门", "+++++++++++++++")
		subordinate.toString()
		for _, employee := range subordinate.subordinates {
			employee.toString()
		}
	}
}
```

结构图

<img src="/img/pattern-composite.png" alt="composite" />  

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/组合模式  
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001      
【菜鸟教程】https://www.runoob.com/design-pattern/composite-pattern.html  