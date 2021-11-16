package 组合模式

import "fmt"

type Employee struct {
	no           int
	name         string
	dept         string
	salary       int
	subordinates []*Employee
}

func NewEmployee(no int, name, dept string, salary int) *Employee {
	return &Employee{
		no:           no,
		name:         name,
		dept:         dept,
		salary:       salary,
		subordinates: []*Employee{},
	}
}

func (e *Employee) Add(em *Employee) {
	e.subordinates = append(e.subordinates, em)
}

func (e *Employee) Remove(em *Employee) {
	for index, item := range e.subordinates {
		if item.no == em.no {
			e.subordinates = e.subordinates[:index+copy(e.subordinates[index:], e.subordinates[index+1:])]
			break
		}
	}
}

func (e *Employee) GetSubordinates() []*Employee {
	return e.subordinates
}

func (e *Employee) ToString() {
	fmt.Println(fmt.Sprintf("Employee:no:%d;name:%s;dept:%s;salary:%d。", e.no, e.name, e.dept, e.salary))
}
