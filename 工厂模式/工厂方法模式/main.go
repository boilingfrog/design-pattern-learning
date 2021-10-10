package main

import "fmt"

func main() {
	apple := apple{}
	fmt.Println(apple.Fruit())

	banana := banana{}
	fmt.Println(banana.Fruit())
}

type apple struct{}

func (*apple) Fruit() string {
	return "我是苹果，我很好吃"
}

type banana struct{}

func (*banana) Fruit() string {
	return "我是香蕉，我最好吃了"
}
