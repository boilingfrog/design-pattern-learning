<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [备忘录模式](#%E5%A4%87%E5%BF%98%E5%BD%95%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 备忘录模式

### 定义

备忘录( Memento ):在不破坏封装性的前提下，获取一个对象的内部状态，并在该对象之处保存该状态。这样以后就可将该对象恢复到原先保存的状态。  

这个模式的意图很明确，主要是用来防丢失、撤销、恢复等。  

### 优点

1、给用户提供了一种可以恢复状态的机制，可以使用户能够比较方便地回到某个历史的状态；  

2、实现了信息的封装，使得用户不需要关心状态的保存细节。  

### 缺点

消耗资源。如果类的成员变量过多，势必会占用比较大的资源，而且每一次保存都会消耗一定的内存。  

### 适用范围

1、需要保存/恢复数据的相关状态场景；  

2、提供一个可回滚的操作。  

### 代码实现

比如我们看视频，比如看了一半，退出观看，下次打开就还能回到观看的位置。   

```go
ype Memento interface{}

type Video struct {
	time int
}

type videoMemento struct {
	time int
}

func (g *Video) Watch(time int) {
	g.time += time
}

func (g *Video) Save() Memento {
	return &videoMemento{
		time: g.time,
	}
}

func (g *Video) Load(m Memento) {
	gm := m.(*videoMemento)
	g.time = gm.time
}

func (g *Video) Status() string {
	return fmt.Sprintf("video watch time:%d", g.time)
}
```

测试代码  

```go
func TestVideo_Watch(t *testing.T) {
	video := &Video{
		time: 10,
	}

	t.Log(video.Status())
	progress := video.Save()

	video.Watch(30)
	t.Log(video.Status())

	video.Load(progress)
	t.Log(video.Status())
}
```

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/备忘录模式      
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001  
【设计模式】https://github.com/senghoo/golang-design-pattern    