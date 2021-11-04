package main

import "sync"

var once sync.Once

func GetOnceInstance() *Tool {
	once.Do(func() {
		instance = &Tool{
			Name: "我sync.once初始化的，我已经初始化了",
		}
	})
	return instance
}
