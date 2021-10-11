<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [工厂模式](#%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F)
  - [简单工厂模式(Simple Factory)](#%E7%AE%80%E5%8D%95%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8Fsimple-factory)
    - [定义](#%E5%AE%9A%E4%B9%89)
    - [优点](#%E4%BC%98%E7%82%B9)
    - [缺点](#%E7%BC%BA%E7%82%B9)
    - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
    - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [工厂方法模式(Factory Method)](#%E5%B7%A5%E5%8E%82%E6%96%B9%E6%B3%95%E6%A8%A1%E5%BC%8Ffactory-method)
    - [定义](#%E5%AE%9A%E4%B9%89-1)
    - [优点](#%E4%BC%98%E7%82%B9-1)
    - [缺点](#%E7%BC%BA%E7%82%B9-1)
    - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4-1)
    - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0-1)
  - [抽象工厂模式(Abstract Factory)](#%E6%8A%BD%E8%B1%A1%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8Fabstract-factory)
    - [定义](#%E5%AE%9A%E4%B9%89-2)
    - [优点](#%E4%BC%98%E7%82%B9-2)
    - [缺点](#%E7%BC%BA%E7%82%B9-2)
    - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4-2)
    - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0-2)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 工厂模式

一般情况下，工厂模式分为三种更加细分的类型：简单工厂、工厂方法和抽象工厂。不过，在GoF的《设计模式》一书中，它将简单工厂模式看作是工厂方法模式的一种特例，所以工厂模式只被分成了工厂方法和抽象工厂两类。  

在这三种细分的工厂模式中，简单工厂、工厂方法原理比较简单，在实际的项目中也比较常用。我们这里来重点的介绍下这两种。  

### 简单工厂模式(Simple Factory)

#### 定义

定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。  

因为在简单工厂模式用于创建实例的方法是静态的方法，因此简单工厂模式又被称为静态工厂方法模式，它属于类创建型模式。  

#### 优点

- 工厂类含有必要的判断逻辑，可以决定在什么时候创建哪一个产品类的实例，客户端可以免除直接创建产品对象的责任，而仅仅“消费”产品；简单工厂模式通过这种做法实现了对责任的分割，它提供了专门的工厂类用于创建对象。

- 客户端无须知道所创建的具体产品类的类名，只需要知道具体产品类所对应的参数即可，对于一些复杂的类名，通过简单工厂模式可以减少使用者的记忆量。

- 通过引入配置文件，可以在不修改任何客户端代码的情况下更换和增加新的具体产品类，在一定程度上提高了系统的灵活性。

#### 缺点

- 由于工厂类集中了所有产品创建逻辑，一旦不能正常工作，整个系统都要受到影响。 

- 使用简单工厂模式将会增加系统中类的个数，在一定程序上增加了系统的复杂度和理解难度。

- 系统扩展困难，一旦添加新产品就不得不修改工厂逻辑，在产品类型较多时，有可能造成工厂逻辑过于复杂，不利于系统的扩展和维护。

- 简单工厂模式由于使用了静态工厂方法，造成工厂角色无法形成基于继承的等级结构。

#### 适用范围 

工厂类负责创建的对象比较少，客户只知道传入了工厂类的参数，对于始何创建对象（逻辑）不关心。  

#### 代码实现

```go
package main

import "fmt"

func main() {
	f := getFruit("apple")
	fmt.Println(f.Fruit())
}

type FruitFactory interface {
	Fruit() string
}

func getFruit(t string) FruitFactory {
	switch t {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}
	}

	return nil
}

type apple struct{}

func (*apple) Fruit() string {
	return "我是苹果，我很好吃"
}

type banana struct{}

func (*banana) Fruit() string {
	return "我是香蕉，我最好吃了"
}
```

总结下：  

主要是通过 if 来判断逻辑，当我们有新的实现需要加入，只需要添加对应的 if 判断就好了。    

UML 类图  

<img src="/img/factory-simple.png" alt="factory" style="zoom:50%;" />

### 工厂方法模式(Factory Method)

#### 定义

工厂方法模式（英语：Factory method pattern）是一种实现了“工厂”概念的面向对象设计模式。就像其他创建型模式一样，它也是处理在不指定对象具体类型的情况下创建对象的问题。工厂方法模式的实质是“定义一个创建对象的接口，但让实现这个接口的类来决定实例化哪个类。工厂方法让类的实例化推迟到子类中进行。”   

#### 优点

- 一个调用者想创建一个对象，只要知道其名称就可以了。  

- 扩展性高，如果想增加一个产品，只要扩展一个工厂类就可以。  

- 屏蔽产品的具体实现，调用者只关心产品的接口。 

#### 缺点

每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。这并不是什么好事。  

#### 适用范围 

当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂。  

#### 代码实现

```go
package main

import "fmt"

func main() {
	apple := apple{}
	fmt.Println(apple.Fruit())

	banana := banana{}
	fmt.Println(banana.Fruit())
}

type FruitFactory interface {
	Fruit() string
}

type apple struct{}

func (*apple) Fruit() string {
	return "我是苹果，我很好吃"
}

type banana struct{}

func (*banana) Fruit() string {
	return "我是香蕉，我最好吃了"
}
```

UML 类图  

<img src="/img/factory-method.png" alt="factory" style="zoom:50%;" />

### 抽象工厂模式(Abstract Factory)

#### 定义

抽象工厂模式（Abstract Factory），提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。  

#### 优点

抽象工厂模式除了具有工厂方法模式的优点外，最主要的优点就是可以在类的内部对产品族进行约束。所谓的产品族，一般或多或少的都存在一定的关联，抽象工厂模式就可以在类内部对产品族的关联关系进行定义和描述，而不必专门引入一个新的类来进行管理。  

#### 缺点

抽象工厂模式在于难于应付“新对象”的需求变动。难以支持新种类的产品。难以扩展抽象工厂以生产新种类的产品。这是因为抽象工厂几乎确定了可以被创建的产品集合，支持新种类的产品就需要扩展该工厂接口，这将涉及抽象工厂类及其所有子类的改变。   

#### 适用范围 

- 1.一个系统不应当依赖于产品类实例如何被创建、组合和表达的细节，这对于所有形态的工厂模式都是重要的。

- 2.这个系统的产品有多于一个的产品族，而系统只消费其中某一族的产品。

- 3.同属于同一个产品族的产品是在一起使用的，这一约束必须在系统的设计中体现出来。（比如：Intel主板必须使用Intel CPU、Intel芯片组）

- 4.系统提供一个产品类的库，所有的产品以同样的接口出现，从而使客户端不依赖于实现。

#### 代码实现

```go
package main

import "fmt"

func main() {
	h := HainanFruit{}
	fmt.Println(h.Apple())
	fmt.Println(h.Banana())

	w := WuhanFruit{}
	fmt.Println(w.Apple())
	fmt.Println(w.Banana())
}

type appleInterface interface {
	Apple()
}

type appleInfo struct {
	Name   string
	Origin string
}

func (appleInfo) Apple() {
	panic("implement me")
}

type bananaInterface interface {
	Banana()
}

type bananaInfo struct {
	Name   string
	Origin string
}

func (bananaInfo) Banana() {
	panic("implement me")
}

type FruitFactory interface {
	Apple() appleInfo
	Banana() bananaInfo
}

type HainanFruit struct{}

func (*HainanFruit) Apple() appleInfo {
	return appleInfo{
		Name:   "我是 苹果；",
		Origin: "产地 海南",
	}
}

func (*HainanFruit) Banana() bananaInfo {
	return bananaInfo{
		Name:   "我是 香蕉；",
		Origin: "产地 海南",
	}
}

type WuhanFruit struct{}

func (*WuhanFruit) Apple() appleInfo {
	return appleInfo{
		Name:   "我是 苹果；",
		Origin: "产地 武汉",
	}
}

func (*WuhanFruit) Banana() bananaInfo {
	return bananaInfo{
		Name:   "我是 香蕉；",
		Origin: "产地 武汉",
	}
}
```

### 参考  

【工厂方法模式】https://wiki.jikexueyuan.com/project/java-design-pattern/factory-pattern.html  
【抽象工厂模式】https://refactoringguru.cn/design-patterns/abstract-factory  
【极客时间】设计模式之美  
【抽象工厂】https://www.liaoxuefeng.com/wiki/1252599548343744/1281319134822433    
【简单工厂模式，工厂方法模式和抽象工厂模式的异同】https://blog.csdn.net/gatieme/article/details/17525805  




  
