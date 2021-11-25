package 访问者模式

import "testing"

func TestVisitor(t *testing.T) {
	var elements []Element
	elements = append(elements, &ConcreteElementA{})
	elements = append(elements, &ConcreteElementB{})

	for _, item := range elements {
		cv1 := &ConcreteVisitor1{}
		cv2 := &ConcreteVisitor2{}
		item.Accept(cv1)
		item.Accept(cv2)
	}
}
