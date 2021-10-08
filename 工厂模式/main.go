package main

import "fmt"

func main() {
	f := newFruit("apple")
	fmt.Println(f.Color())
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

//type apple struct{}
//
//func (*apple) Color() string {
//	return "我是苹果，我是青色的"
//}
//
//type banana struct{}
//
//func (*banana) Color() string {
//	return "我是香蕉，我是黄色的"
//}
