<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [适配器模式](#%E9%80%82%E9%85%8D%E5%99%A8%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代理、桥接、装饰器、适配器4种设计模式的区别](#%E4%BB%A3%E7%90%86%E6%A1%A5%E6%8E%A5%E8%A3%85%E9%A5%B0%E5%99%A8%E9%80%82%E9%85%8D%E5%99%A84%E7%A7%8D%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F%E7%9A%84%E5%8C%BA%E5%88%AB)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 适配器模式

### 定义

适配器模式的英文翻译是Adapter Design Pattern。顾名思义，这个模式就是用来做适配的，它将不兼容的接口转换为可兼容的接口，让原本由于接口不兼容而不能一起工作的类可以一起工作。  

举个栗子：  

现在比较新款的电脑都有USB-C接口，但是我们目前的鼠标键盘的接口都是传统的USB接口，所以是不能使用的，这时候我们会买个转接口来进行接口的转接，那么这个转接口在设计模式中就是适配器。   

### 代码实现

```
// 基础的播放功能
type MediaPlayer interface {
	play(audioType string, fileName string)
}

// 不同的播放器平台
type AdvancedMediaPlayer interface {
	playVlc(fileName string)
	playMp4(fileName string)
}

// VlcPlayers
type VlcPlayers struct {
}

func (v *VlcPlayers) playVlc(fileName string) {
	fmt.Println("正在播放" + fileName)
}

func (v *VlcPlayers) playMp4(fileName string) {
	fmt.Println("格式不支持")
}

// Mp4Player
type Mp4Player struct {
}

func (m *Mp4Player) playVlc(fileName string) {
	fmt.Println("格式不支持")
}

func (m *Mp4Player) playMp4(fileName string) {
	fmt.Println("正在播放" + fileName)
}

// 适配器
type MediaAdapter struct {
	MusicPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	var mediaAdapter MediaAdapter
	switch audioType {
	case "vlc":
		mediaAdapter.MusicPlayer = &VlcPlayers{}
	case "mp4":
		mediaAdapter.MusicPlayer = &Mp4Player{}
	default:
		panic("不支持的类型")
	}
	return &mediaAdapter
}
func (m *MediaAdapter) play(audioType string, fileName string) {
	switch audioType {
	case "vlc":
		m.MusicPlayer.playVlc(fileName)
	case "mp4":
		m.MusicPlayer.playMp4(fileName)

	}
}

// AudioPlayer 音频播放器类
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

// Play 播放音频
func (auPlayer *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3" {
		fmt.Println("正在播放" + fileName)
		return
	}
	auPlayer.mediaAdapter = NewMediaAdapter(audioType)
	auPlayer.mediaAdapter.play(audioType, fileName)
}
```

测试文件

```go
func TestPlayer(t *testing.T) {
	ad := AudioPlayer{}
	ad.Play("mp4", "荷塘月色")
	ad.Play("vlc", "小苹果")
	ad.Play("mp3", "天空之城")
}
```

这里做个简单的分析  

1、我们有一个 AudioPlayer ，但是只能播放 mp3;  

2、我们希望 AudioPlayer 也可以播放 mp3 和 vlc;  

3、引入了一个 MediaAdapter ，通过适配器来处理不支持的功能，对于 AudioPlayer 来讲，它只用需要调用 MediaAdapter 就能实现各种播放格式音频的播放；  

4、MediaAdapter 对各种格式进行了包装，不同的格式音频，可以有用相同的调用方法。 

放一张结构图  

<img src="/img/pattern-adapter.png" alt="adapter" />

### 优点

1、可以让任何两个没有关联的类一起运行。 

2、提高了类的复用。 

3、增加了类的透明度。 

4、灵活性好。

### 缺点

过多地使用适配器，会让系统非常零乱，不易整体进行把握   

一般来说，适配器模式可以看作一种“补偿模式”，用来补救设计上的缺陷。应用这种模式算是“无奈之举”，如果在设计初期，我们就能协调规避接口不兼容的问题，那这种模式就没有应用的机会了。  

如果大量的使用这种模式，可能就是我们的前期的设计有很大的问题，就需要考虑重构了  

### 适用范围

1、封装有缺陷的接口设计  

2、统一多个类的接口设计  

3、替换依赖的外部系统  

4、兼容老版本接口  

5、适配不同格式的数据  

### 代理、桥接、装饰器、适配器4种设计模式的区别  

**代理模式**：代理模式在不改变原始类接口的条件下，为原始类定义一个代理类，主要目的是控制访问，而非加强功能，这是它跟装饰器模式最大的不同。  

**桥接模式**：桥接模式的目的是将接口部分和实现部分分离，从而让它们可以较为容易、也相对独立地加以改变。  

**装饰器模式**：装饰者模式在不改变原始类接口的情况下，对原始类功能进行增强，并且支持多个装饰器的嵌套使用。  

**适配器模式**：适配器模式是一种事后的补救策略。适配器提供跟原始类不同的接口，而代理模式、装饰器模式提供的都是跟原始类相同的接口。  

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/%E8%A3%85%E9%A5%B0%E5%99%A8%E6%A8%A1%E5%BC%8F  
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001    
【菜鸟教程】https://www.runoob.com/design-pattern/adapter-pattern.html  