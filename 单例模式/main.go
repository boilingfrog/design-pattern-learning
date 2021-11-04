package main

import (
	"fmt"
	"sync"
)

// 使用结构体代替类
type Tool struct {
	Name string
}

// 锁对象
var lock sync.Mutex

// 建立私有变量
var instance *Tool

// 加锁保证线程安全
func GetInstance() *Tool {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &Tool{
			Name: "我已经初始化了",
		}
	}
	return instance
}

func main() {
	fmt.Println(GetInstance().Name)
}
