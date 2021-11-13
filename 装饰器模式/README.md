<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [装饰器模式](#%E8%A3%85%E9%A5%B0%E5%99%A8%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 装饰器模式

### 定义

装饰模式：动态的给一些对象添加额外的职责，就增加功能来说，装饰模式比生成子类更灵活。  

举个栗子：  

我们现在买手机或者电脑，都有基础配置，然后根据我们选择的运行内存的大小，具体的CPU，手机或电脑的价格最终价格就是不一样的。这里就用到装饰模式，定制的内存和CPU对我们的设备进行了装饰作用。   

看下结构图  

<img src="/img/pattern-decorator.png" alt="decorator" />

### 代码实现

```go
// 基础款 16g运存 10代cpu
type basicPhone struct {
}

func (p *basicPhone) getPrice() int {
	return 2000
}

// 32g运存
type choose32RAMPhone struct {
	phone phone
}

func (r *choose32RAMPhone) getPrice() int {
	price := r.phone.getPrice()
	return price + 500
}

// 11代CPU
type choose11CPUPhone struct {
	phone phone
}

func (r *choose11CPUPhone) getPrice() int {
	price := r.phone.getPrice()
	return price + 1000
}
```

测试代码

```go
func TestDecorator(t *testing.T) {
	phone := &basicPhone{}

	choose32RAMPhone := &choose32RAMPhone{
		phone: phone,
	}

	orderPhone := &choose11CPUPhone{
		phone: choose32RAMPhone,
	}

	t.Log(orderPhone.getPrice())
}
```

### 优点

1、装饰器是继承的有力补充，比继承灵活，在不改变原有对象的情况下，动态的给一个对象扩展功能，即插即用  

2、通过使用不用装饰类及这些装饰类的排列组合，可以实现不同效果  

3、装饰器模式完全遵守开闭原则  

装饰器模式主要解决继承关系过于复杂的问题，通过组合来替代继承。它主要的作用是给原始类添加增强功能。这也是判断是否该用装饰器模式的一个重要的依据。除此之外，装饰器模式还有一个特点，那就是可以对原始类嵌套使用多个装饰器。为了满足这个应用场景，在设计的时候，装饰器类需要跟原始类继承相同的抽象类或者接口。  

### 缺点

装饰器模式会增加许多子类，过度使用会增加程序得复杂性。  

### 适用范围

1、扩展一个类的功能。   

2、动态增加功能，动态撤销。  

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/%E6%A1%A5%E6%8E%A5%E6%A8%A1%E5%BC%8F    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001  
