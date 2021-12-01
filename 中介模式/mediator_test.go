package 中介模式

import "testing"

func TestMediator_ForwardMessage(t *testing.T) {
	// 创建中介者，联合国
	mediator := &UnitedNationsSecurityCouncil{}

	usa := USA{mediator}
	mediator.USA = usa
	usa.SendMess("不准研制核武器，否则要发动战争了")

	irap := Irap{mediator}
	mediator.Irap = irap
	irap.SendMess("我们没有核武器")
}
