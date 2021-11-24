<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [迭代器模式](#%E8%BF%AD%E4%BB%A3%E5%99%A8%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 迭代器模式

### 定义

迭代器模式（Iterator Design Pattern），也叫作游标模式（Cursor Design Pattern）。  

提供了一种方法顺序的访问一个聚合对象中的各个元素，而不是暴露该对象的内部表示。   

这里的聚合对象也叫容器聚合对象，实质上及时包含一组对象的对象，例如数组、链表、树、图、跳表。迭代器模式将集合对象的遍历操作从集合类中拆分出来，放到迭代器类中，让两者的职责更加单一。   

一个通俗的总结对于迭代器模式：  

流水线上坐一天，每个包裹扫一遍。   

### 优点

迭代器相比于 for 循环的优点  

1、迭代器模式封装集合内部的复杂数据结构，开发者不需要了解如何遍历，直接使用容器提供的迭代器即可；  

2、迭代器模式将集合对象的遍历操作从集合类中拆分出来，放到迭代器类中，让两者的职责更加单一；  

3、迭代器模式让添加新的遍历算法更加容易，更符合开闭原则。除此之外，因为迭代器都实现自相同的接口，在开发中，基于接口而非实现编程，替换迭代器也变得更加容易。  

### 缺点

由于迭代器模式将存储数据和遍历数据的职责分离，增加新的聚合类需要对应增加新的迭代器类，类的个数成对增加，这在一定程度上增加了系统的复杂性。  

### 适用范围

1、访问一个聚合对象的内容而无须暴露它的内部表示；  

2、需要为聚合对象提供多种遍历方式；  

3、为遍历不同的聚合结构提供一个统一的接口。  

### 代码实现

使用迭代器输出切片中的名字集合   

```go
type Iterator interface {
	HasNext() bool
	Next() string
}

type names []string

func (na names) NewIterator() *NameRepository {
	return &NameRepository{
		index: 0,
		names: na,
	}
}

type NameRepository struct {
	index int
	names names
}

func (nr *NameRepository) HasNext() bool {
	if nr.index < len(nr.names) {
		return true
	}
	return false
}

func (nr *NameRepository) Next() string {
	if nr.HasNext() {
		name := nr.names[nr.index]
		nr.index++
		return name
	}

	return ""
}
```

测试代码  

```go
func TestIterator(t *testing.T) {
	names := names{
		"小明", "小豆", "小龙",
	}
	nameRepository := names.NewIterator()
	for nameRepository.HasNext() {
		t.Log(nameRepository.Next())
	}
}
```

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/迭代器模式   
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   