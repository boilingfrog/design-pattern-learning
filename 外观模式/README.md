<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [外观模式](#%E5%A4%96%E8%A7%82%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [关于接口粒度的思考](#%E5%85%B3%E4%BA%8E%E6%8E%A5%E5%8F%A3%E7%B2%92%E5%BA%A6%E7%9A%84%E6%80%9D%E8%80%83)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 外观模式

### 定义

外观模式也叫门面模式  

外观模式(Facade),为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。   

### 适用范围

**1、解决易用性问题**    

门面模式可以用来封装系统的底层实现，隐藏系统的复杂性，提供一组更加简单易用、更高层的接口。  

**2、解决性能问题**  

我们通过将多个接口调用替换为一个门面接口调用，减少网络通信成本，提高App客户端的响应速度。  

假设有一个系统A，提供了a、b、c、d四个接口。系统B完成某个业务功能，需要调用A系统的a、b、d接口。利用门面模式，我们提供一个包裹a、b、d接口调用的门面接口x，给系统B直接使用。  

**3、解决分布式事务问题**

这个直接来个栗子吧  

比如我们现在设计微服务，一个用户服务，一个金币服务。每个服务都提供对外的增删查改等操作。现在我们有一个需求，新用户登陆之后送用户金币。简单地调用，肯定是创建一个用户信息，之后调用金币服务给这个用户加金币。当然实际地开发中我们肯定要考虑分布式事务，保障这两个操作肯定能一起成功或失败，不能出现一个成功一个失败的场景。当然我们常规的做法肯定是引入分布式框架，或者补偿的机制来处理。  

其实借助于门面模式的思想也是可以处理，我们可以设计一个包裹这两个操作的新接口，让新接口在一个事务中执行两个SQL操作。   

### 代码实现

假设有一个系统A，提供了a、b、c、d四个接口。系统B完成某个业务功能，需要调用A系统的a、b、d接口。利用门面模式，我们提供一个包裹a、b、d接口调用的门面接口x，给系统B直接使用。    

```go
type User struct {
}

func (u *User) GetUser(userId int) {
	fmt.Println("获取用户的信息")
}

type GoldCoin struct {
}

func (g *GoldCoin) GetUserGoldCoin(userId int) {
	fmt.Println("获取用户金币的信息")
}

type Order struct {
}

func (o *Order) GetUserOrder(userId int) {
	fmt.Println("获取用户订单信息")
}

func GetUserInfo(userId int) {
	user := User{}
	user.GetUser(userId)

	goldCoin := GoldCoin{}
	goldCoin.GetUserGoldCoin(userId)

	order := Order{}
	order.GetUserOrder(userId)
}
```

放一张结构图  

<img src="/img/pattern-facade.png" alt="facade" />

### 优点

- 对客户屏蔽子系统组件，减少了客户处理的对象数目并使得子系统使用起来更加容易。  

- 实现了子系统与客户之间的松耦合关系，这使得子系统的组件变化不会影响到调用它的客户类，只需要调整外观类即可。  

- 降低了大型软件系统中的编译依赖性，并简化了系统在不同平台之间的移植过程，因为编译一个子系统一般不需要编译所有其他的子系统。一个子系统的修改对其他子系统没有任何影响，而且子系统内部变化也不会影响到外观对象。  

- 只是提供了一个访问子系统的统一入口，并不影响用户直接使用子系统类。  

### 缺点

- 不能很好地限制客户使用子系统类，如果对客户访问子系统类做太多的限制则减少了可变性和灵活性。  

- 在不引入抽象外观类的情况下，增加新的子系统可能需要修改外观类或客户端的源代码，违背了“开闭原则”。 

### 关于接口粒度的思考  

接口粒度设计得太大，太小都不好。太大会导致接口不可复用，太小会导致接口不易用。在实际的开发中，接口的可复用性和易用性需要“微妙”的权衡。  

针对这个问题，王争大佬给出了一个原则，可以作为参考，**尽量保持接口的可复用性，但针对特殊情况，允许提供冗余的门面接口，来提供更易用的接口**。     

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/%E5%A4%96%E8%A7%82%E6%A8%A1%E5%BC%8F  
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   
【外观模式】https://design-patterns.readthedocs.io/zh_CN/latest/structural_patterns/facade.html  
