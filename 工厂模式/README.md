<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [工厂模式](#%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F)
  - [简单工厂模式(Simple Factory)](#%E7%AE%80%E5%8D%95%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8Fsimple-factory)
    - [定义](#%E5%AE%9A%E4%B9%89)
    - [优点](#%E4%BC%98%E7%82%B9)
    - [缺点](#%E7%BC%BA%E7%82%B9)
    - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [工厂方法模式(Factory Method)](#%E5%B7%A5%E5%8E%82%E6%96%B9%E6%B3%95%E6%A8%A1%E5%BC%8Ffactory-method)
    - [定义](#%E5%AE%9A%E4%B9%89-1)
    - [优点](#%E4%BC%98%E7%82%B9-1)
    - [缺点](#%E7%BC%BA%E7%82%B9-1)
    - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0-1)

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

#### 代码实现

```go
package main

import "fmt"

func main() {
	f := newFruit("apple")
	fmt.Println(f.Color())
}

type Fruit interface {
	Color() string
}

func newFruit(t string) Fruit {
	switch t {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}
	}

	return nil
}

type apple struct{}

func (*apple) Color() string {
	return "我是苹果，我是青色的"
}

type banana struct{}

func (*banana) Color() string {
	return "我是香蕉，我是黄色的"
}
```

总结下：  

主要是通过 if 来判断逻辑，当我们有新的实现需要加入，只需要添加对应的 if 判断就好了。    

UML 类图  

<img src="/img/factory-simple.png" alt="factory" style="zoom:50%;" />

### 工厂方法模式(Factory Method)

#### 定义

工厂方法模式（英语：Factory method pattern）是一种实现了“工厂”概念的面向对象设计模式。就像其他创建型模式一样，它也是处理在不指定对象具体类型的情况下创建对象的问题。工厂方法模式的实质是“定义一个创建对象的接口，但让实现这个接口的类来决定实例化哪个类。工厂方法让类的实例化推迟到子类中进行。”   

当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂。  

#### 优点

- 一个调用者想创建一个对象，只要知道其名称就可以了。  

- 扩展性高，如果想增加一个产品，只要扩展一个工厂类就可以。  

- 屏蔽产品的具体实现，调用者只关心产品的接口。 

#### 缺点

每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。这并不是什么好事。  

#### 代码实现

```go
package main

import "fmt"

func main() {
	apple := apple{}
	fmt.Println(apple.Color())

	banana := banana{}
	fmt.Println(banana.Color())
}

type apple struct{}

func (*apple) Color() string {
	return "我是苹果，我是青色的"
}

type banana struct{}

func (*banana) Color() string {
	return "我是香蕉，我是黄色的"
}
```

### 抽象工厂模式(Abstract Factory)

#### 定义

抽象工厂模式（Abstract Factory），提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。  

#### 优点

#### 缺点




  
