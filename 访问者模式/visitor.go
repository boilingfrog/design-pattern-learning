package 访问者模式

import "fmt"

type Visitor interface {
	VisitConcreteElementA(cea *ConcreteElementA)
	VisitConcreteElementB(ceb *ConcreteElementB)
}

type ConcreteVisitor1 struct {
}

func (cea *ConcreteVisitor1) VisitConcreteElementA(concreteElementA *ConcreteElementA) {
	fmt.Println("concreteVisitor1 visitConcreteElementA")
}

func (*ConcreteVisitor1) VisitConcreteElementB(concreteElementB *ConcreteElementB) {
	fmt.Println("concreteVisitor1 visitConcreteElementB")
}

type ConcreteVisitor2 struct {
}

func (*ConcreteVisitor2) VisitConcreteElementA(concreteElementA *ConcreteElementA) {
	fmt.Println("concreteVisitor2 visitConcreteElementA")
}

func (*ConcreteVisitor2) VisitConcreteElementB(concreteElementB *ConcreteElementB) {
	fmt.Println("concreteVisitor2 visitConcreteElementB")
}

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct {
}

func (cea *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(cea)
}

type ConcreteElementB struct {
}

func (ceb *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ceb)
}
