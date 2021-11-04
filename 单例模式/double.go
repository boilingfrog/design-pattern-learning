package main

//第一次判断不加锁，第二次加锁保证线程安全，一旦对象建立后，获取对象就不用加锁了。
func GetDoubleInstance() *Tool {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &Tool{
				Name: "我是双重检测，我已经初始化了",
			}
		}
		lock.Unlock()
	}
	return instance
}
