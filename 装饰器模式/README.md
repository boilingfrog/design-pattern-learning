## 装饰器模式

### 定义

装饰模式：动态的给一些对象添加额外的职责，就增加功能来说，装饰模式比生成子类更灵活。  

举个栗子：  

我们现在买手机或者电脑，都有基础配置，然后根据我们选择的运行内存的大小，具体的CPU，手机或电脑的价格最终价格就是不一样的。这里就用到装饰模式，定制的内存和CPU对我们的设备进行了装饰作用。   




### 代码实现

```go
// 基础款 16g运存 10代cpu
type basicPhone struct {
}

func (p *basicPhone) getPrice() int {
	return 2000
}

// 32g运存
type choose32RAMPhone struct {
	phone phone
}

func (r *choose32RAMPhone) getPrice() int {
	price := r.phone.getPrice()
	return price + 500
}

// 11代CPU
type choose11CPUPhone struct {
	phone phone
}

func (r *choose11CPUPhone) getPrice() int {
	price := r.phone.getPrice()
	return price + 1000
}
```

测试代码

```go
func TestDecorator(t *testing.T) {
	phone := &basicPhone{}

	choose32RAMPhone := &choose32RAMPhone{
		phone: phone,
	}

	orderPhone := &choose11CPUPhone{
		phone: choose32RAMPhone,
	}

	t.Log(orderPhone.getPrice())
}
```

### 优点

### 缺点

### 适用范围



### 参考

