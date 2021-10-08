package main

type Fruit interface {
	Color() string
}

type apple struct{}

func (*apple) Color() string {
	return "我是苹果，我是青色的"
}

type banana struct{}

func (*banana) Color() string {
	return "我是香蕉，我是黄色的"
}
