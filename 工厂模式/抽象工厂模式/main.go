package main

import "fmt"

func main() {
	h := HainanFruit{}
	fmt.Println(h.Apple())

	w := WuhanFruit{}
	fmt.Println(w.Apple())
}

type FruitFactory interface {
	Apple() string
	Banana() string
}

type HainanFruit struct{}

func (*HainanFruit) Apple() string {
	return "我是海南的苹果"
}

func (*HainanFruit) Banana() string {
	return "我是海南的香蕉"
}

type WuhanFruit struct{}

func (*WuhanFruit) Apple() string {
	return "我是武汉的苹果"
}

func (*WuhanFruit) Banana() string {
	return "我是武汉的香蕉"
}
