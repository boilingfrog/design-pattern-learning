<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [解释器模式](#%E8%A7%A3%E9%87%8A%E5%99%A8%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 解释器模式

### 定义

解释器模式(interpreter):给定一种语言，定义它的文法的一种表示，并定一个解释器，这个解释器使用该表示来解释语言中的句子。  

解释器模式的意义在于，它分离多种复杂功能的实现，每个功能只需关注自身的解释。  

对于调用者不用关心内部的解释器的工作，只需要用简单的方式组合命令就可以。  

### 优点

1、可扩展性比较好，灵活;  

2、增加了新的解释表达式的方式;   

3、易于实现简单文法。   

### 缺点

1、可利用场景比较少;   

2、对于复杂的文法比较难维护;   

3、解释器模式会引起类膨胀。   

### 适用范围

解释器模式的代码实现比较灵活，没有固定的模板。我们前面也说过，应用设计模式主要是应对代码的复杂性，实际上，解释器模式也不例外。它的代码实现的核心思想，就是将语法解析的工作拆分到各个小类中，以此来避免大而全的解析类。一般的做法是，将语法规则拆分成一些小的独立的单元，然后对每个单元进行解析，最终合并为对整个语法规则的解析。    

### 代码实现

这里简单实现了一个加减的运算器，我们对每种运算定义对应的方法，避免所有的运算操作放到一个函数中，这就体现了解释器模式的核心思想，将语法解析的工作拆分到各个小类中，以此来避免大而全的解析类。  

```go
type Expression interface {
	Interpret() int
}

type NumberExpression struct {
	val int
}

func (n *NumberExpression) Interpret() int {
	return n.val
}

type AdditionExpression struct {
	left, right Expression
}

func (n *AdditionExpression) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

type SubtractionExpression struct {
	left, right Expression
}

func (n *SubtractionExpression) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Expression
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAdditionExpression()
		case "-":
			p.prev = p.newSubtractionExpression()
		default:
			p.prev = p.newNumberExpression()
		}
	}
}

func (p *Parser) newAdditionExpression() Expression {
	p.index++
	return &AdditionExpression{
		left:  p.prev,
		right: p.newNumberExpression(),
	}
}

func (p *Parser) newSubtractionExpression() Expression {
	p.index++
	return &SubtractionExpression{
		left:  p.prev,
		right: p.newNumberExpression(),
	}
}

func (p *Parser) newNumberExpression() Expression {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &NumberExpression{
		val: v,
	}
}

func (p *Parser) Result() Expression {
	return p.prev
}
```

测试代码  

```go
func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 3 + 3 + 3 + 3")
	res := p.Result().Interpret()
	expect := 13
	if res != expect {
		t.Fatalf("expect %d got %d", expect, res)
	}
	t.Log(res)
}
```

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/解释器模式    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001  
【设计模式】https://github.com/senghoo/golang-design-pattern    