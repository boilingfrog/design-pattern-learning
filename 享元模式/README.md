<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [享元模式](#%E4%BA%AB%E5%85%83%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用场景](#%E9%80%82%E7%94%A8%E5%9C%BA%E6%99%AF)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [享元模式和单例模式的区别](#%E4%BA%AB%E5%85%83%E6%A8%A1%E5%BC%8F%E5%92%8C%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F%E7%9A%84%E5%8C%BA%E5%88%AB)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 享元模式

### 定义

享元模式(Flyweight),运用共享技术有效的支持大量细粒度的对象。  

享元模式的意图是复用对象，节省内存，前提是享元对象是不可变对象。就是当一个系统中有大量的重复对象的时候，如果这些对象是不可变对象，我们就可以使用
享元模式，将这些对象设计成享元，在内存只保存一份，供需要的代码使用，这样能减少内存中对象的数量，起到节省内存的作用。实际上，不仅仅相同对象可以设计成享元，对于相似对象，我们也可以将这些对象中相同的部分（字段）提取出来，设计成享元，让这些大量相似对象引用这些享元。   

不可变对象是函数初始化之后，他的状态不会改变了，也就是不会存在被修改的情况。  

### 优点

大大减少对象的创建，降低系统的内存，使效率提高。  

### 缺点

提高了系统的复杂度   

### 适用场景

一个系统中有大量相同或者相似的对象，使用这些对象，会造成内存的大量消耗，这时候可以考虑使用享元模式  

### 代码实现

我们都下过象棋，对于象棋一共有32枚棋子，每方各16枚。每次下棋使用的都是这32枚棋子，只是每次移动的棋子的位置。   

那么这32枚棋子就可以作为享元，每次下棋就不用重新初始化生成棋子。只需在棋盘中放置棋子的位子，下棋的时候，更改位置即可。      

```go
// ChessPieceUnit ...
type ChessPieceUnit struct {
	id    int
	text  string
	color string
}

var pieces map[int]*ChessPieceUnit

func init() {
	pieces = make(map[int]*ChessPieceUnit, 32)
	pieces[1] = &ChessPieceUnit{
		id:    1,
		text:  "马",
		color: "BLACK",
	}
	pieces[2] = &ChessPieceUnit{
		id:    2,
		text:  "炮",
		color: "BLACK",
	}
	// ...
}

func getChessPiece(chessPieceId int) *ChessPieceUnit {
	return pieces[chessPieceId]
}

type ChessPiece struct {
	chessPieceUnit *ChessPieceUnit
	positionX      int
	positionY      int
}

func newChessPiece(chessPieceId int, positionX, positionY int) *ChessPiece {
	return &ChessPiece{
		chessPieceUnit: getChessPiece(chessPieceId),
		positionX:      positionX,
		positionY:      positionY,
	}
}

// 棋盘
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func (cb *ChessBoard) InitChessBoard() {
	cb.chessPieces = make(map[int]*ChessPiece, 32)
	cb.chessPieces[1] = newChessPiece(1, 0, 1)
	cb.chessPieces[2] = newChessPiece(2, 0, 2)
	// ...
}

// Move 下棋
func (cb *ChessBoard) Move(chessPieceId int, positionX, positionY int) {
	// TODO
}
```

### 享元模式和单例模式的区别

单例模式：一个类只能创建一个对象。  

享元模式：一个类可以创建多个对象，每个对象被多处代码引用共享。实际上，享元模式有点类似于之前讲到的单例的变体：多例。  

单例对象可以是可变的。 享元对象是不可变的。  

实现上差不多，不过设计意图不同，享元模式是为了对象复用，节省内存，而应用多例模式是为了限制对象的个数。  

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/享元模式  
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001     
【享元模式】https://boilingfrog.github.io/2021/11/17/%E4%BD%BF%E7%94%A8go%E5%AE%9E%E7%8E%B0%E4%BA%AB%E5%85%83%E6%A8%A1%E5%BC%8F/  