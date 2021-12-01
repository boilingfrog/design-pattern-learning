package 中介模式

import (
	"fmt"
)

type Country interface {
	SendMess(message string)
	GetMess(message string)
}

type USA struct {
	mediator *UnitedNationsSecurityCouncil
}

func (usa *USA) SendMess(message string) {
	usa.mediator.ForwardMessage(usa, message)
}

func (usa *USA) GetMess(message string) {
	fmt.Println("USA 获得对方的消息：", message)
}

type Irap struct {
	mediator *UnitedNationsSecurityCouncil
}

func (ir *Irap) SendMess(message string) {
	ir.mediator.ForwardMessage(ir, message)
}

func (ir *Irap) GetMess(message string) {
	fmt.Println("Irap 获得对方的消息：", message)
}

type Mediator interface {
	ForwardMessage(country Country, message string)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Irap
}

func (uns *UnitedNationsSecurityCouncil) ForwardMessage(country Country, message string) {
	switch country.(type) {
	case *USA:
		uns.Irap.GetMess(message)
	case *Irap:
		uns.USA.GetMess(message)
	default:
		fmt.Println("国家不在联合国")
	}
}
