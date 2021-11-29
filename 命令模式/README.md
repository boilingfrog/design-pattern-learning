## 命令模式

### 定义

命令模式(Command)：将一个请求封装成一个对象，从而是你可用不同的的请求对客户进行参数化；对请求排队或记录请求日志，以及支持可撤销的操作。  

### 优点

1、降低了系统耦合度；  

2、新的命令可以很容易添加到系统中去。   

### 缺点

使用命令模式可能会导致某些系统有过多的具体命令类。   

### 适用范围

1、如果你需要通过操作来参数化对象， 可使用命令模式。  

2、如果你想要将操作放入队列中、 操作的执行或者远程执行操作， 可使用命令模式。   

3、如果你想要实现操作回滚功能， 可使用命令模式。  

命令模式的主要作用和应用场景，是用来控制命令的执行，比如，异步、延迟、排队执行命令、撤销重做命令、存储命令、给命令记录日志等等，这才是命令模式能发挥独一无二作用的地方。  

### 代码实现

```go
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
```

测试文件  

```go
func TestNewCommand(t *testing.T) {
	r := Receiver{}
	concreteCommand := NewConcreteCommand(r)

	invoker := Invoker{}
	invoker.SetCommand(concreteCommand)
	invoker.ExecuteCommand()
}
```

结构图  

<img src="/img/pattern-command.png" alt="command" />

### 命令模式对比策略模式

对于一些相似的设计模式的区分，我们应该关注设计意图，应用场景的不同，而不是只看解决方案这一部分，或者只关注代码实现。单看实现，一些设计模式确实很相似，比较难区分。    

从设计意图区分：  

策略模式：不同的策略具有相同的目的、不同的实现、互相之间可以替换。  

命令模式：不同的命令具有不同的目的，对应不同的处理逻辑，并且互相之间不可替换。    

### 参考

【文中代码】https://github.com/boilingfrog/design-pattern-learning/tree/master/命令模式       
【大话设计模式】https://book.douban.com/subject/2334288/  
【极客时间】https://time.geekbang.org/column/intro/100039001   