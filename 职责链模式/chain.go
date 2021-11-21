package 职责链模式

import "fmt"

type Teacher interface {
	HaveRight(day int) bool
	HandleApproveRequest(name string, day int) bool
}

type RequestChain struct {
	Teacher
	approver *RequestChain
}

func (r *RequestChain) SetApprover(m *RequestChain) {
	r.approver = m
}

func (r *RequestChain) HandleApproveRequest(name string, day int) bool {
	if r.Teacher.HaveRight(day) {
		return r.Teacher.HandleApproveRequest(name, day)
	}
	if r.approver != nil {
		return r.approver.HandleApproveRequest(name, day)
	}
	fmt.Println("请假时间太久了，不予审批")
	return false
}

func (r *RequestChain) HaveRight(day int) bool {
	return true
}

type HeadTeacher struct{}

func NewHeadTeacherChain() *RequestChain {
	return &RequestChain{
		Teacher: &HeadTeacher{},
	}
}

func (*HeadTeacher) HaveRight(day int) bool {
	return day <= 2
}

func (*HeadTeacher) HandleApproveRequest(name string, day int) bool {
	fmt.Println(fmt.Sprintf("班主任，批准了%s的请假申请,请假天数%d", name, day))
	return true
}

type DepTeacher struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Teacher: &DepTeacher{},
	}
}

func (*DepTeacher) HaveRight(day int) bool {
	return day <= 7
}

func (*DepTeacher) HandleApproveRequest(name string, day int) bool {
	fmt.Println(fmt.Sprintf("系主任，批准了%s的请假申请,请假天数%d", name, day))
	return true
}

type DeanTeacher struct{}

func NewDeanTeacherChain() *RequestChain {
	return &RequestChain{
		Teacher: &DeanTeacher{},
	}
}

func (*DeanTeacher) HaveRight(day int) bool {
	return day <= 10
}

func (*DeanTeacher) HandleApproveRequest(name string, day int) bool {
	fmt.Println(fmt.Sprintf("院长，批准了%s的请假申请,请假天数%d", name, day))
	return true
}
