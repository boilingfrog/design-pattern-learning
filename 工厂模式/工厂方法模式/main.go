package main

import "fmt"

func main() {
	apple := appleFactory{}
	fmt.Println(apple.Fruit())

	banana := bananaFactory{}
	fmt.Println(banana.Fruit())
}

type Fruit interface {
	Fruit() string
}

type appleFactory struct{}

func (*appleFactory) Fruit() string {
	return "我是苹果，我很好吃"
}

type bananaFactory struct{}

func (*bananaFactory) Fruit() string {
	return "我是香蕉，我最好吃了"
}
