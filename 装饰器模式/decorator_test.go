package 装饰器模式

import (
	"testing"
)

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
