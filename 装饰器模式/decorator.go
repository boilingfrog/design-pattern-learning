package 装饰器模式

type phone interface {
	getPrice() int
}

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
