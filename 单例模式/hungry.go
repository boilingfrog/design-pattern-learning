package main

var cfg *config

func init() {
	cfg = &config{
		Name: "我被初始化了",
	}
}

type config struct {
	Name string
}

// NewConfig 提供获取实例的方法
func NewConfig() *config {
	return cfg
}
