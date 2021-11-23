<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [状态模式](#%E7%8A%B6%E6%80%81%E6%A8%A1%E5%BC%8F)
  - [定义](#%E5%AE%9A%E4%B9%89)
  - [优点](#%E4%BC%98%E7%82%B9)
  - [缺点](#%E7%BC%BA%E7%82%B9)
  - [适用范围](#%E9%80%82%E7%94%A8%E8%8C%83%E5%9B%B4)
  - [代码实现](#%E4%BB%A3%E7%A0%81%E5%AE%9E%E7%8E%B0)
  - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 状态模式

### 定义

状态模式(state):当一个条件的内在状态改变时允许改变其行为，这个对象看起来像是改变了其类。   

状态模式主要解决的是当控制一个对象状态转换的的条件表达式过于复杂，把状态的判断逻辑转移到表示不同状态的一些列类当中，可以把复杂的判断逻辑简化。   

如果状态判断很简单，就没必要使用状态模式了。   

### 状态机

状态机有3个组成部分：状态（State）、事件（Event）、动作（Action）。  

状态机的实现方式有三种：  

1、分支逻辑法； 

就是参照状态转移图，将每一个状态转移，原模原样地直译成代码。这样编写的代码会包含大量的if-else或switch-case分支判断逻辑，甚至是嵌套的分支判断逻辑。  

2、查表法；  

就是把对应的状态，以及状态对应的时间放入到数据或者 map 中，这样通过根据业务的状态，查询数据或者 map 就能找到对应需要执行的事件。   

对于状态很多、状态转移比较复杂的状态机来说，查表法比较合适。通过二维数组来表示状态转移图，能极大地提高代码的可读性和可维护性。  

3、状态模式。  

如果状态对应的执行操作很复杂，需要一些列的复杂的逻辑操作，这时候就需要引入状态模式了。  

### 优点

1、将与特定状态相关的代码放在单独的类中；  

2、无需修改已有状态类和上下文就能引入新状态；  

3、通过消除臃肿的状态机条件语句简化上下文代码。   

### 缺点

1、状态模式的使用必然会增加系统类和对象的个数；  

2、如果状态机只有几个状态，或者很少发生改变，使用状态模式反而会使系统变的复杂。   

### 适用范围

1、如果对象需要根据自身当前状态进行不同行为，同时状态的数量非常多且与状态相关的代码会频繁变更的话，可使用状态模式；  

2、如果某个类需要根据成员变量的当前值改变自身行为，从而需要使用大量的条件语句时，可使用该模式；  

3、当相似状态和基于条件的状态机转换中存在许多重复代码时，可使用状态模式。    

### 代码实现

栗子：  

比如公司员工，一般会有小组长，经理，总经理等职务，每个职务处理的工作内容，拥有的权限责任是不同的，不同的岗位的职责就可以通过状态模式来表示。   

```go
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
```

测试代码  

```go
func TestCompanyEmployee(t *testing.T) {
	cem := InitCompanyEmployeeMachine()
	fmt.Println("员工的职级", cem.GetRankState(), ";员工的薪水", cem.Salary(), ";员工的待遇", cem.HaveRight())
	// 晋级
	cem.Promotion()
	fmt.Println("员工的职级", cem.GetRankState(), ";员工的薪水", cem.Salary(), ";员工的待遇", cem.HaveRight())
}
```

输出的内容  

```
员工的职级 employee ;员工的薪水 1000 ;员工的待遇 [休假，年终奖]
员工的职级 groupLeader ;员工的薪水 2000 ;员工的待遇 [休假，年终奖 绩效奖金]
```

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/状态模式    
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   