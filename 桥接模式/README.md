<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [桥接模式](#%E6%A1%A5%E6%8E%A5%E6%A8%A1%E5%BC%8F)
  - [前言](#%E5%89%8D%E8%A8%80)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [应用场景](#%E5%BA%94%E7%94%A8%E5%9C%BA%E6%99%AF)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 桥接模式

### 前言

桥接模式的代码实现非常简单，但是理解起来稍微有点难度，并且应用场景也比较局限，所以，相当于代理模式来说，桥接模式在实际的项目中并没有那么常用，所以能分辨出来就行了，这里不做重点的学习。   

### 定义

桥接模式：将抽象部分和它的实现部分分离，使他们都可以独立的变化。   

看完定义还是一脸懵😳 

举个栗子吧：使用大话模式中的手机栗子来分析    

我们知道不同品牌的手机以及不同品牌手机的不同版本对于手机软件的要求是不一样的，可能在M品牌中的软件，就不能在N品牌的手机中使用。M品牌中10代版本的软件可能就不能在1代版本中安装使用。  
   
我们如果使用代码去实现这种关系：  

1、按照品牌分类实现的结构图

<img src="/img/bridge-brand.png" alt="bridge" />

2、按照软件分类实现的结构图

<img src="/img/bridge-soft.png" alt="bridge"  />

上面的第一种和第二种实现方式  

如果增加在增加品牌手机和软件，那么修改就是灾难级别的，原来已经写好的模块也需要做修改    

3、使用桥接模式实现的结构图

<img src="/img/bridge.png" alt="bridge" />

第三种也就是我们讲的桥接模式，如果有手机品牌和软件类型的加入，只需要进行扩展就好了，之前已经写好的模块就不用修改了。  

由于实现的方式是多种的，桥接模式的核心就是把这些实现独立出来，让他们自己变化。这样每种变化不会影响到其他的实现，从而达到应对变化的目的。   

### 优点

1、抽象与实现分离，扩展能力强  

2、符合开闭原则  

3、符合合成复用原则  

4、其实现细节对客户透明   

通俗点讲就是我们的实现系统可能有多角度的分类，每种类都有可能变化，桥架模式做的就是把这些多角度独立出来，让他们可以自己变化，而不影响其他的模块，减少他们之间的藕合。   

### 缺点

桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系建立在抽象层，要求开发者针对抽象进行设计与编程。   

### 应用场景

 1、如果一个系统需要在构件的抽象化角色和具体化角色之间增加更多的灵活性，避免在两个层次之间建立静态的继承联系，通过桥接模式可以使它们在抽象层建立一个关联关系。   
 
 2、对于那些不希望使用继承或因为多层次继承导致系统类的个数急剧增加的系统，桥接模式尤为适用。   
 
 3、一个类存在两个独立变化的维度，且这两个维度都需要进行扩展。   

### 代码实现

还是上面手机的栗子  

```go
// 手机软件
type HandsetSoft interface {
	Run() string
}

// 手机游戏
type HandsetGame struct {
}

func (hg *HandsetGame) Run() string {
	return "运行手机游戏"
}

// 手机通讯录
type HandsetAddressList struct {
}

func (hg *HandsetAddressList) Run() string {
	return "运行手机通讯录"
}

// 手机品牌
type HandsetBrand interface {
	SetHandsetSoft(HandsetSoft)
}

// M品牌手机手机
type HandsetBrandM struct {
	HandsetSoft
}

func NewHandsetBrandM() *HandsetBrandM {
	return &HandsetBrandM{}
}

func (hw *HandsetBrandM) SetHandsetSoft(soft HandsetSoft) {
	hw.HandsetSoft = soft
}

func (hw *HandsetBrandM) Run() string {
	return "M品牌的手机-" + hw.HandsetSoft.Run()
}

// N品牌的手机
type HandsetBrandN struct {
	HandsetSoft
}

func NewHandsetBrandN() *HandsetBrandN {
	return &HandsetBrandN{}
}

func (ap *HandsetBrandN) SetHandsetSoft(soft HandsetSoft) {
	ap.HandsetSoft = soft
}

func (ap *HandsetBrandN) Run() string {
	return "N品牌的手机-" + ap.HandsetSoft.Run()
}
```

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/%E6%A1%A5%E6%8E%A5%E6%A8%A1%E5%BC%8F    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001  