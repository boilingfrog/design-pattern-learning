package 状态模式

import (
	"fmt"
	"testing"
)

func TestCompanyEmployee(t *testing.T) {
	cem := InitCompanyEmployeeMachine()
	fmt.Println("员工的职级", cem.GetRankState(), ";员工的薪水", cem.Salary(), ";员工的待遇", cem.HaveRight())
	// 晋级
	cem.Promotion()
	fmt.Println("员工的职级", cem.GetRankState(), ";员工的薪水", cem.Salary(), ";员工的待遇", cem.HaveRight())
}
