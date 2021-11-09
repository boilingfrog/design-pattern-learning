package 代理模式

import "fmt"

type GiveGift interface {
	GiveDolls() string
	GiveFlowers() string
	GiveChocolate() string
}

// 追求者
type Pursuit struct {
	GirlName string
}

func NewGirl(name string) *Pursuit {
	return &Pursuit{
		GirlName: name,
	}
}

func (ps *Pursuit) GiveDolls() string {
	return fmt.Sprintf("%s-送你娃娃", ps.GirlName)
}

func (ps *Pursuit) GiveFlowers() string {
	return fmt.Sprintf("%s-送你漂亮的鲜花", ps.GirlName)
}

func (ps *Pursuit) GiveChocolate() string {
	return fmt.Sprintf("%s-送你巧克力", ps.GirlName)
}

// 代理也就是中间人
type Proxy struct {
	Pursuit
}

func NewProxy(name string) *Pursuit {
	return NewGirl(name)
}

func (pr *Proxy) GiveDolls() string {
	return pr.GiveDolls()
}

func (pr *Proxy) GiveFlowers() string {
	return pr.GiveFlowers()
}

func (pr *Proxy) GiveChocolate() string {
	return pr.GiveChocolate()
}
