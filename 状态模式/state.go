package 状态模式

type RankState string

const (
	Employee       RankState = "employee"       // 员工
	GroupLeader    RankState = "groupLeader"    // 小组长
	Manager        RankState = "manager"        // 经理
	GeneralManager RankState = "generalManager" // 总经理
)

type CompanyState interface {
	getRankState() RankState
	salary(*CompanyEmployeeMachine) int
	haveRight(*CompanyEmployeeMachine) []string
	promotion(*CompanyEmployeeMachine)
}

// CompanyEmployee
type CompanyEmployee struct {
}

func (ce *CompanyEmployee) getRankState() RankState {
	return Employee
}

func (ce *CompanyEmployee) salary(cm *CompanyEmployeeMachine) int {
	return cm.salary
}

func (ce *CompanyEmployee) promotion(cm *CompanyEmployeeMachine) {
	cm.salary += 1000
	cm.right = append(cm.right, "绩效奖金")
	cm.companyState = &CompanyGroupLeader{}
}

func (ce *CompanyEmployee) haveRight(cm *CompanyEmployeeMachine) []string {
	return cm.right
}

// CompanyGroupLeader
type CompanyGroupLeader struct {
}

func (ce *CompanyGroupLeader) getRankState() RankState {
	return GroupLeader
}

func (ce *CompanyGroupLeader) salary(cm *CompanyEmployeeMachine) int {
	return cm.salary
}

func (ce *CompanyGroupLeader) promotion(cm *CompanyEmployeeMachine) {
	cm.salary += 1000
	cm.companyState = &CompanyGroupLeader{}
}

func (ce *CompanyGroupLeader) haveRight(cm *CompanyEmployeeMachine) []string {
	return cm.right
}

type CompanyEmployeeMachine struct {
	salary       int
	right        []string
	companyState CompanyState
}

func InitCompanyEmployeeMachine() *CompanyEmployeeMachine {
	return &CompanyEmployeeMachine{
		salary: 1000,
		right: []string{
			"休假，年终奖",
		},
		companyState: &CompanyEmployee{},
	}
}

func (m *CompanyEmployeeMachine) HaveRight() []string {
	return m.companyState.haveRight(m)
}

func (m *CompanyEmployeeMachine) Salary() int {
	return m.companyState.salary(m)
}

func (m *CompanyEmployeeMachine) GetRankState() RankState {
	return m.companyState.getRankState()
}

func (m *CompanyEmployeeMachine) Promotion() {
	m.companyState.promotion(m)
}
