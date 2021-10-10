package main

import "fmt"

func main() {
	f := newFruit("apple")
	fmt.Println(f.Fruit())
}

type FruitFactory interface {
	Fruit() string
}

func newFruit(t string) FruitFactory {
	switch t {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}
	}

	return nil
}

type apple struct{}

func (*apple) Fruit() string {
	return "我是苹果，我很好吃"
}

type banana struct{}

func (*banana) Fruit() string {
	return "我是香蕉，我最好吃了"
}
