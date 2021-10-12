package main

import "fmt"

func main() {
	f := WuhanFruitFactory{}
	b := f.ChooseApple()
	b.Fruit()
}

type FruitInterface interface {
	ChooseApple() ProductInterface
	ChooseBanana() ProductInterface
}

type ProductInterface interface {
	Fruit()
}

type HainanApple struct {
}

func (h HainanApple) Fruit() {
	fmt.Println("我是苹果，来自海南")
}

type HainanBanana struct {
}

func (h HainanBanana) Fruit() {
	fmt.Println("我是香蕉，来自海南")
}

type WuhanApple struct {
}

func (w WuhanApple) Fruit() {
	fmt.Println("我是苹果，来自武汉")
}

type WuhanBanana struct {
}

func (w WuhanBanana) Fruit() {
	fmt.Println("我是香蕉，来自武汉")
}

type WuhanFruitFactory struct {
}

func (w WuhanFruitFactory) ChooseApple() ProductInterface {
	return WuhanApple{}
}

func (w WuhanFruitFactory) ChooseBanana() ProductInterface {
	return WuhanBanana{}
}

type HainanFruitFactory struct {
}

func (gd HainanFruitFactory) ChooseApple() ProductInterface {
	return HainanApple{}
}

func (gd HainanFruitFactory) ChooseBanana() ProductInterface {
	return HainanBanana{}
}
