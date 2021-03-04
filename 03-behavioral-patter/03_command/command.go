package command

import "fmt"

// 执行命令的接口
type Command interface {
	Execute()
}

// 命令包裹的对象实例
type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

// 封装主板的 启动命令实例
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

// 封装主板的, 重启命令实例
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

// 主机,绑定命令,实现调度
type Box struct {
	buttion1 Command
	buttion2 Command
}

func NewBox(buttion1, buttion2 Command) *Box {
	return &Box{
		buttion1: buttion1,
		buttion2: buttion2,
	}
}

func (b *Box) PressButtion1() {
	b.buttion1.Execute()
}

func (b *Box) PressButtion2() {
	b.buttion2.Execute()
}
