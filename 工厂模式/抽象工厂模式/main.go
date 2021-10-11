package main

import "fmt"

func main() {
	h := HainanFruit{}
	fmt.Println(h.Apple())

	w := WuhanFruit{}
	fmt.Println(w.Apple())
}

type fruit struct {
	Name   string
	Origin string
}

type FruitFactory interface {
	Apple() fruit
	Banana() fruit
}

type HainanFruit struct{}

func (*HainanFruit) Apple() fruit {
	return fruit{
		Name:   "我是 苹果；",
		Origin: "产地 海南",
	}
}

func (*HainanFruit) Banana() fruit {
	return fruit{
		Name:   "我是 香蕉；",
		Origin: "产地 海南",
	}
}

type WuhanFruit struct{}

func (*WuhanFruit) Apple() fruit {
	return fruit{
		Name:   "我是 苹果；",
		Origin: "产地 武汉",
	}
}

func (*WuhanFruit) Banana() fruit {
	return fruit{
		Name:   "我是 香蕉；",
		Origin: "产地 武汉",
	}
}
