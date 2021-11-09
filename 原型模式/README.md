<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [原型模式](#%E5%8E%9F%E5%9E%8B%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用场景](#%E9%80%82%E7%94%A8%E5%9C%BA%E6%99%AF)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 原型模式

### 定义

如果对象的创建成本比较大，而同一个类的不同对象之间差别不大（大部分字段都相同），在这种情况下，我们可以利用对已有对象（原型）进行复制（或者叫拷贝）的方式来创建新对象，以达到节省创建时间的目的。这种基于原型来创建对象的方式就叫作原型设计模式（Prototype Design Pattern），简称原型模式。  

原型模式是能基于拷贝来的，对于拷贝我们知道有两种形式，深拷贝和浅拷贝  

浅拷贝只复制指向某个对象的指针，而不复制对象本身，新旧对象还是共享同一块内存。但深拷贝会另外创造一个一模一样的对象，新对象跟原对象不共享内存，修改新对象不会改到原对象。  

原型模式浅拷贝：

1、省内存，拷贝时间更快；    

2、浅拷贝容易出现原始数据被修改的情况，一般不建议使用；  

3、浅拷贝可以拷贝不可变对象；  

原型模式深拷贝：

1、数据完全隔离；  

2、不过数据量大的情况下，深拷贝比起浅拷贝来说，更加耗时，更加耗内存空间；   

### 代码实现

```go
// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
```

测试文件  

```go
var (
	deepCopyManager    *PrototypeManager
	shallowCopyManager *PrototypeManager
)

// 深拷贝实现Cloneable
type DeepCopy struct {
	name string
}

func (t *DeepCopy) Clone() Cloneable {
	tc := *t
	return &tc
}

// 浅拷贝实现Cloneable
type ShallowCopy struct {
	name string
}

func (t *ShallowCopy) Clone() Cloneable {
	return t
}

func TestDeepCopyClone(t *testing.T) {
	t1 := deepCopyManager.Get("dc")

	t2 := t1.Clone()
	// 深拷贝，指向的不是同一个变量的地址
	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}

	t21 := t2.(*DeepCopy)
	t21.name = "ShallowCopy-test"

	t11 := t1.(*DeepCopy)
	// 深拷贝name，不会影响到copy前的变量
	if t11.name == t21.name {
		t.Fatal("shallowCopy err")
	}
}

func TestShallowCopyClone(t *testing.T) {
	t1 := shallowCopyManager.Get("sc")

	t2 := t1.Clone()
	// 浅拷贝，变量地址的指向不变
	if t1 != t2 {
		t.Fatal("error! get clone not working")
	}

	t21 := t2.(*ShallowCopy)
	t21.name = "ShallowCopy-test"

	t11 := t1.(*ShallowCopy)
	// 深拷贝name，copy之前的变量和copy之后的变量同时更改
	if t11.name != t21.name {
		t.Fatal("shallowCopy err")
	}
}

func init() {
	deepCopyManager = NewPrototypeManager()

	dc := &DeepCopy{
		name: "deepCopy",
	}
	deepCopyManager.Set("dc", dc)

	shallowCopyManager = NewPrototypeManager()
	sc := &ShallowCopy{
		name: "shallowCopy",
	}
	shallowCopyManager.Set("sc", sc)
}
```

### 优点

1、使用原型模式创建对象比直接new一个对象在性能上要好的多，因为是直接进行的内存拷贝，比初始化性能上会好很多；   

2、简化对象的创建，对于创建对象就像我们在编辑文档时的复制粘贴一样简单。

### 缺点

克隆包含循环引用的复杂对象可能会非常麻烦。  

### 适用场景

1、在项目中，如果存在大量相同或相似对象的创建，如果用传统的构造函数来创建对象，会比较复杂和耗费资源，用原型模式生产对象就很高效；  

2、对象创建过程比较麻烦，但复制比较简单的时候；   

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/%E5%8E%9F%E5%9E%8B%E6%A8%A1%E5%BC%8F    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   
【原型模式】https://github.com/senghoo/golang-design-pattern    
【原文地址】https://boilingfrog.github.io/2021/11/08/%E4%BD%BF%E7%94%A8go%E5%AE%9E%E7%8E%B0%E5%8E%9F%E5%9E%8B%E6%A8%A1%E5%BC%8F/    


