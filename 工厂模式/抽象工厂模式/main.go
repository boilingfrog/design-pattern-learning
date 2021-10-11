package main

import "fmt"

func main() {
	h := HainanFruit{}
	fmt.Println(h.Apple())
	fmt.Println(h.Banana())

	w := WuhanFruit{}
	fmt.Println(w.Apple())
	fmt.Println(w.Banana())
}

type appleInterface interface {
	Apple()
}

type appleInfo struct {
	Name   string
	Origin string
}

func (appleInfo) Apple() {
	panic("implement me")
}

type bananaInterface interface {
	Banana()
}

type bananaInfo struct {
	Name   string
	Origin string
}

func (bananaInfo) Banana() {
	panic("implement me")
}

type FruitFactory interface {
	Apple() appleInfo
	Banana() bananaInfo
}

type HainanFruit struct{}

func (*HainanFruit) Apple() appleInfo {
	return appleInfo{
		Name:   "我是 苹果；",
		Origin: "产地 海南",
	}
}

func (*HainanFruit) Banana() bananaInfo {
	return bananaInfo{
		Name:   "我是 香蕉；",
		Origin: "产地 海南",
	}
}

type WuhanFruit struct{}

func (*WuhanFruit) Apple() appleInfo {
	return appleInfo{
		Name:   "我是 苹果；",
		Origin: "产地 武汉",
	}
}

func (*WuhanFruit) Banana() bananaInfo {
	return bananaInfo{
		Name:   "我是 香蕉；",
		Origin: "产地 武汉",
	}
}
