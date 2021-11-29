package 命令模式

import "testing"

func TestNewCommand(t *testing.T) {
	r := Receiver{}
	concreteCommand := NewConcreteCommand(r)

	invoker := Invoker{}
	invoker.SetCommand(concreteCommand)
	invoker.ExecuteCommand()
}
