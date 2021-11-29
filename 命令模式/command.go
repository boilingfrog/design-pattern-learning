package 命令模式

import "fmt"

type Receiver struct {
}

func (*Receiver) Action() {
	fmt.Println("执行命令")
}

type Command struct {
	receiver Receiver
}

func NewCommand(receiver Receiver) *Command {
	return &Command{
		receiver: receiver,
	}
}

type CommandImpl interface {
	Execute()
}

type ConcreteCommand struct {
	*Command
}

func NewConcreteCommand(receiver Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		NewCommand(receiver),
	}
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

type Invoker struct {
	command CommandImpl
}

func (ik *Invoker) ExecuteCommand() {
	ik.command.Execute()
}

func (ik *Invoker) SetCommand(command CommandImpl) {
	ik.command = command
}
