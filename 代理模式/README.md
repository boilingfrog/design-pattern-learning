<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [代理模式](#%E4%BB%A3%E7%90%86%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [应用场景](#%E5%BA%94%E7%94%A8%E5%9C%BA%E6%99%AF)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 代理模式

### 定义

定义：为其对象提供一种代理以控制这个对象的访问。通俗点讲就是它在不改变原始类（或叫被代理类）代码的情况下，通过引入代理类来给原始类附加功能。   

来点通俗的理解：   

比如我们买火车票，除了火车站，很多代售点也是可以买的，代售点的作用就是代理模式   

### 优点

1、代理模式在客户端与目标对象之间起到一个中介作用和保护目标对象的作用；  

2、代理对象可以扩展目标对象的功能；  

3、代理模式能将客户端与目标对象分离，在一定程度上降低了系统的耦合度，增加了程序的可扩展性；  

### 缺点

1、代理模式会造成系统设计中类的数量增加；    

2、在客户端和目标对象之间增加一个代理对象，会造成请求处理速度变慢；  

3、增加了系统的复杂度；

### 应用场景

- 1、业务系统的非功能性需求开发  

一些非功能性的业务需求，比如：监控、统计、鉴权、限流、事务、幂等、日志。我们将这些附加功能与业务功能解耦，放到代理类中统一处理，让程序员只需要关注业务方面的开发。   

- 2、代理模式在RPC中的应用  

RPC框架也可以看成是一种代理模式。   

GoF的《设计模式》一书中把它称作远程代理。通过远程代理，将网络通信、数据编解码等细节隐藏起来。客户端使用RPC服务就像使用本地函数一样，RPC服务的开发者也只需要开发业务逻辑，就像开发本地使用的函数一样，不需要关注跟客户端的交互细节。  

- 3、代理模式在缓存中的应用  

对于一些接口的开发，有时候对于一个功能，我们会提供两种接口，一种支持缓存，一种不支持缓存，对于需要实时数据的需求，我们让其调用实时查询接口，对于不需要实时数据的需求，我们让其调用支持缓存的接口。  

简单的做法就是接口写两个，一个支持缓存的一个不支持缓存，但是这样我们的代码就有些臃肿了。  

可以使用代理模式中的动态代理。举个栗子：  

如果是基于Spring框架来开发的话，那就可以在AOP切面中完成接口缓存的功能。在应用启动的时候，我们从配置文件中加载需要支持缓存的接口，以及相应的缓存策略（比如过期时间）等。当请求到来的时候，我们在AOP切面中拦截请求，如果请求中带有支持缓存的字段（比如http://…?..&cached=true），我们便从缓存（内存缓存或者Redis缓存等）中获取数据直接返回。  

### 代码实现

这里借助于大话设计模式中的追女孩的栗子：   

小明喜欢小红，但是害羞不好意思出面，所以拜托好朋友，小张，出面给小红送礼物。   

小明是追求者，小张就是小明的中间人，也就是代理。    

```go
type GiveGift interface {
	GiveDolls() string
	GiveFlowers() string
	GiveChocolate() string
}

// 追求者
type Pursuit struct {
	GirlName string
}

func NewGirl(name string) *Pursuit {
	return &Pursuit{
		GirlName: name,
	}
}

func (ps *Pursuit) GiveDolls() string {
	return fmt.Sprintf("%s-送你娃娃", ps.GirlName)
}

func (ps *Pursuit) GiveFlowers() string {
	return fmt.Sprintf("%s-送你漂亮的鲜花", ps.GirlName)
}

func (ps *Pursuit) GiveChocolate() string {
	return fmt.Sprintf("%s-送你巧克力", ps.GirlName)
}

// 代理也就是中间人
type Proxy struct {
	Pursuit
}

func NewProxy(name string) *Pursuit {
	return NewGirl(name)
}

func (pr *Proxy) GiveDolls() string {
	return pr.GiveDolls()
}

func (pr *Proxy) GiveFlowers() string {
	return pr.GiveFlowers()
}

func (pr *Proxy) GiveChocolate() string {
	return pr.GiveChocolate()
}
```

最后放上一张结构图  

<img src="/img/pattern-proxy.png" alt="proxy" />

### 参考
 
【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/%E4%BB%A3%E7%90%86%E6%A8%A1%E5%BC%8F    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001    
【代理模式】https://boilingfrog.github.io/2021/11/08/%E4%BD%BF%E7%94%A8go%E5%AE%9E%E7%8E%B0%E4%BB%A3%E7%90%86%E6%A8%A1%E5%BC%8F/  
