package main

import "fmt"

func main() {
	apple := apple{}
	fmt.Println(apple.Color())

	banana := banana{}
	fmt.Println(banana.Color())
}

type apple struct{}

func (*apple) Color() string {
	return "我是苹果，我是青色的"
}

type banana struct{}

func (*banana) Color() string {
	return "我是香蕉，我是黄色的"
}

type Fruit interface {
}

func newFruit(t string) Fruit {
	switch t {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}
	}

	return nil
}
