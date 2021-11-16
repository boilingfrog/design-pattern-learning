package 组合模式

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	ceo := NewEmployee(1, "马云", "CEO", 1000000)

	headMarketing := NewEmployee(2, "小白", "市场总监", 50000)
	clerk1 := NewEmployee(1001, "小明", "Marketing", 2000)
	clerk2 := NewEmployee(1002, "小张", "Marketing", 1000)

	cto := NewEmployee(3, "小龙", "CTO", 90000)
	tc1 := NewEmployee(1003, "马龙", "Technology", 3000)
	tc2 := NewEmployee(1004, "张龙", "Technology", 4000)

	ceo.Add(headMarketing)
	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)

	ceo.Add(cto)
	cto.Add(tc1)
	cto.Add(tc2)

	ceo.ToString()
	for _, subordinate := range ceo.subordinates {
		fmt.Println("部门", "+++++++++++++++")
		subordinate.ToString()
		for _, employee := range subordinate.subordinates {
			employee.ToString()
		}
	}
}
