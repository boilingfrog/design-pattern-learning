<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [单例模式](#%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
    - [懒汉模式](#%E6%87%92%E6%B1%89%E6%A8%A1%E5%BC%8F)
    - [饿汉模式](#%E9%A5%BF%E6%B1%89%E6%A8%A1%E5%BC%8F)
    - [双重检测](#%E5%8F%8C%E9%87%8D%E6%A3%80%E6%B5%8B)
    - [sync.Once](#synconce)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 单例模式

### 定义

什么是单例模式：保证一个类仅有一个实例，并提供一个全局访问它的全局访问点。      

例如：在某个服务器程序中，该服务器的配置信息存放在一个文件中，这些配置数据由一个单例对象统一读取，然后服务进程中的其他对象再通过这个单例对象获取这些配置信息。这样方便了读取，同时保证了我们的配置信息只会初始化一次。   

### 优点

1、在单例模式中，活动的单例只有一个实例，对单例类的所有实例化得到的都是相同的一个实例。这样就 防止其它对象对自己的实例化，确保所有的对象都访问一个实例   

2、单例模式具有一定的伸缩性，类自己来控制实例化进程，类就在改变实例化进程上有相应的伸缩性     

3、提供了对唯一实例的受控访问  

4、由于在系统内存中只存在一个对象，因此可以节约系统资源，当需要频繁创建和销毁的对象时单例模式无疑可以提高系统的性能  

5、允许可变数目的实例  

6、避免对共享资源的多重占用  

### 缺点

1、不适用于变化的对象，如果同一类型的对象总是要在不同的用例场景发生变化，单例就会引起数据的错误，不能保存彼此的状态  

2、由于单利模式中没有抽象层，因此单例类的扩展有很大的困难  

3、单例类的职责过重，在一定程度上违背了“单一职责原则”  

4、滥用单例将带来一些负面问题，如为了节省资源将数据库连接池对象设计为的单例类，可能会导致共享连接池对象的程序过多而出现连接池溢出；如果实例化的对象长时间不被利用，系统会认为是垃圾而被回收，这将导致对象状态的丢失  

### 适用范围

我们在项目中使用单例，都是用它来表示一些全局唯一类，比如：配置信息类、连接池类、ID生成器类。  

### 代码实现

#### 懒汉模式

懒汉模式也就是在需要的时候，才去创建实例对象。懒汉式相对于饿汉式的优势是支持延迟加载。不过懒汉模式不是线程安全的，需要加锁。  

```go
// 使用结构体代替类
type Tool struct {
	Name string
}

// 锁对象
var lock sync.Mutex

// 建立私有变量
var instance *Tool

// 加锁保证线程安全
func GetInstance() *Tool {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &Tool{
			Name: "我已经初始化了",
		}
	}
	return instance
}
```

懒汉式的缺点也很明显，我们给getInstance()这个方法加了一把锁，导致这个函数的并发度很低。量化一下的话，并发度是1，也就相当于串行操作了。而这个函数是在单例使用期间，一直会被调用。如果这个单例类偶尔会被用到，那这种实现方式还可以接受。但是，如果频繁地用到，那频繁加锁、释放锁及并发度低等问题，会导致性能瓶颈，这种实现方式就不可取了。

#### 饿汉模式

饿汉模式的实现方式比较简单。在类加载的时候，instance静态实例就已经创建并初始化好了，所以，instance实例的创建过程是线程安全的。  

```go
var cfg *config

func init() {
	cfg = &config{
		Name: "我被初始化了",
	}
}

type config struct {
	Name string
}

// NewConfig 提供获取实例的方法
func NewConfig() *config {
	return cfg
}
```

这种方式就是资源提前初始化，有人会讲需要的时候才去初始化，能够避免资源的浪费，不过也有不同的看法。  

1、如果初始化耗时长，那我们最好不要等到真正要用它的时候，才去执行这个耗时长的初始化过程，这会影响到系统的性能（比如，在响应客户端接口请求的时候，做这个初始化操作，会导致此请求的响应时间变长，甚至超时）。采用饿汉式实现方式，将耗时的初始化操作，提前到程序启动的时候完成，这样就能避免在程序运行的时候，再去初始化导致的性能问题。  

2、如果实例占用资源多，按照fail-fast的设计原则（有问题及早暴露），那我们也希望在程序启动时就将这个实例初始化好。如果资源不够，就会在程序启动的时候触发报错（比如Java中的 PermGen Space OOM），我们可以立即去修复。这样也能避免在程序运行一段时间后，突然因为初始化这个实例占用资源过多，导致系统崩溃，影响系统的可用性。  

#### 双重检测

饿汉式不支持延迟加载，懒汉式有性能问题，不支持高并发。这里又引入了一种双重检测的方法。  

在这种实现方式中，只要instance被创建之后，即便再调用getInstance()函数也不会再进入到加锁逻辑中了。  

来看下代码的实现  

```go
// 使用结构体代替类
type Tool struct {
	Name string
}

//锁对象
var lock sync.Mutex

var instance *Tool

//第一次判断不加锁，第二次加锁保证线程安全，一旦对象建立后，获取对象就不用加锁了。
func GetInstance() *Tool {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &Tool{
				Name: "我是双重检测，我已经初始化了",
			}
		}
		lock.Unlock()
	}
	return instance
}
```

#### sync.Once

go 中也提供了 sync.Once 这个方法，来控制只执行一次，具体源码参见[go中sync.Once源码解读](https://www.cnblogs.com/ricklz/p/14503674.html)  

```go
// 使用结构体代替类
type Tool struct {
	Name string
}

var instance *Tool

var once sync.Once

func GetOnceInstance() *Tool {
	once.Do(func() {
		instance = &Tool{
			Name: "我sync.once初始化的，我已经初始化了",
		}
	})
	return instance
}
```

### 参考

【单例模式】https://zh.wikipedia.org/wiki/%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001    
【单例模式的优缺点和使用场景】https://www.cnblogs.com/damsoft/p/6105122.html    