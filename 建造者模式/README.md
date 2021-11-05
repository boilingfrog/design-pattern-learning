<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [建造者模式](#%E5%BB%BA%E9%80%A0%E8%80%85%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 建造者模式

### 定义

建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。  

一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。  

看了定义还是晕晕的，这里来个栗子  

这里按照[设计模式之美](https://time.geekbang.org/column/intro/100039001)中的资源池的例子来进行讨论  

假设有这样一道设计面试题：我们需要定义一个资源池配置类ResourcePoolConfig。这里的资源池，你可以简单理解为线程池、连接池、对象池等。在这个资源池配置类中，有以下几个成员变量，也就是可配置项。现在，请你编写代码实现这个ResourcePoolConfig类。  

| 成员变量 | 解释          | 是否必填 | 默认值 |
| ------  |   -------    | ------ |------ |
| name    | 资源名称       | 是     | 没有 |
| maxTotal| 最大资源数量    | 否     | 10 |
| maxIdle | 最大空闲资源数量 | 否     | 10 |
| minIdle | 最小空闲资源数量 | 否     | 0 |

很简单，来看下代码实现   

```go
const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 10
	defaultMinIdle  = 1
)

// ResourcePoolConfig ...
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func NewResourcePoolConfig(name string, maxTotal, maxIdle, minIdle *int) (*ResourcePoolConfig, error) {
	rc := &ResourcePoolConfig{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}
	if name == "" {
		return nil, errors.New("name is empty")
	}
	rc.name = name

	if maxTotal != nil {
		if *maxTotal <= 0 {
			return nil, errors.New("maxTotal should be positive")
		}
		rc.maxTotal = *maxTotal
	}

	if maxIdle != nil {
		if *maxIdle <= 0 {
			return nil, errors.New("maxIdle should not be negative")
		}
		rc.maxIdle = *maxIdle
	}

	if minIdle != nil {
		if *minIdle <= 0 {
			return nil, errors.New("minIdle should not be negative")
		}
		rc.minIdle = *minIdle
	}

	return rc, nil
}
```

我们接着讨论，如果需要传入的参数过多，我们可以使用 set() 函数来给成员变量赋值，以替代冗长的构造函数。  

```go
func NewResourcePoolConfigSet(name string) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	return &ResourcePoolConfig{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
		name:     name,
	}, nil
}

// SetMinIdle ...
func (rc *ResourcePoolConfig) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("min idle cannot < 0, input: %d", minIdle)
	}
	rc.minIdle = minIdle
	return nil
}

// SetMaxIdle ...
func (rc *ResourcePoolConfig) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max idle cannot < 0, input: %d", maxIdle)
	}
	rc.maxIdle = maxIdle
	return nil
}

// SetMaxTotal ...
func (rc *ResourcePoolConfig) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max total cannot <= 0, input: %d", maxTotal)
	}
	rc.maxTotal = maxTotal
	return nil
}
```

到这里，我们还是没有用上建造者模式，我们来接着分析上面的栗子  

1、上面的 name 字段是必填的，如果必填字段很多，那么我们的函数中又会出现参数很长的情况。当然必填项是不能放在set中设置的，因为如果对应的set没加，我们不能判断该参数必填的逻辑。   

2、比如依赖关系，比如，如果用户设置了maxTotal、maxIdle、minIdle其中一个，就必须显式地设置另外两个；或者配置项之间有一定的约束条件，比如，maxIdle和minIdle要小于等于maxTotal。所以我们就需要一开始就知道所有的参数，才能进行对应校验。  

3、如果我们希望ResourcePoolConfig类对象是不可变对象，也就是说，对象在创建好之后，就不能再修改内部的属性值。要实现这个功能，我们就不能在ResourcePoolConfig类中暴露set()方法。  

这时候建造者模式就登场了  

```go
// ResourcePoolConfig ...
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// ResourcePoolConfigBuilder ...
type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// SetName ...
func (rb *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name can not be empty")
	}
	rb.name = name
	return nil
}

// SetMinIdle ...
func (rb *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("max total cannot < 0, input: %d", minIdle)
	}
	rb.minIdle = minIdle
	return nil
}

// SetMaxIdle ...
func (rb *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max total cannot < 0, input: %d", maxIdle)
	}
	rb.maxIdle = maxIdle
	return nil
}

// SetMaxTotal ...
func (rb *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max total cannot <= 0, input: %d", maxTotal)
	}
	rb.maxTotal = maxTotal
	return nil
}

// Build ...
func (rb *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if rb.name == "" {
		return nil, errors.New("name can not be empty")
	}

	// 设置默认值
	if rb.minIdle == 0 {
		rb.minIdle = defaultMinIdle
	}

	if rb.maxIdle == 0 {
		rb.maxIdle = defaultMaxIdle
	}

	if rb.maxTotal == 0 {
		rb.maxTotal = defaultMaxTotal
	}

	if rb.maxTotal < rb.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", rb.maxTotal, rb.maxIdle)
	}

	if rb.minIdle > rb.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", rb.maxIdle, rb.minIdle)
	}

	return &ResourcePoolConfig{
		name:     rb.name,
		maxTotal: rb.maxTotal,
		maxIdle:  rb.maxIdle,
		minIdle:  rb.minIdle,
	}, nil
}
```

建造者模式，避免了无效状态的存在，因为是设置构建者的变量，构建的变量符合条件之后，一次性的创建对象，这样创建的对象就一直处于有效状态了。  

### 适用范围



### 优点


### 缺点





### 代码实现

### 参考

【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   
【建造者模式】https://www.runoob.com/design-pattern/builder-pattern.html    
【Go设计模式03-建造者模式】https://lailin.xyz/post/builder.html  
