package 观察者模式

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	subject := NewSubject()
	reader1 := NewCustomer("小明")
	reader2 := NewCustomer("小红")
	reader3 := NewCustomer("小李")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	for i := 1; i <= 10; i++ {
		subject.UpdateContext(fmt.Sprintf("更新了%d", i))
		fmt.Println("+++++++++++++++++++++++++++++++++")
	}
}
