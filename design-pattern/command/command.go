package main

import "fmt"

// 命令模式： 是一种数据驱动的设计模式，它属于行为型模式。请求以命令的形式包裹在对象中，并传给调用对象。
// 调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象，该对象执行命令。

// 命令执行者
type Dog struct {
	Name   string
	Action string
}

func (d *Dog) Sit() {
	fmt.Println(d.Name, d.Action)
}

func (d *Dog) Handshake() {
	fmt.Println(d.Name, d.Action)
}

func (d *Dog) Run() {
	fmt.Println(d.Name, d.Action)
}

func NewDog(name, action string) *Dog {
	return &Dog{Name: name, Action: action}
}

//  命令接口
type Command interface {
	Execute()
}

// -------- 命令接口实现 -----------
type SitCommand struct {
	dog Dog
}

func (c *SitCommand) Execute() {
	c.dog.Sit()
}

type HandshakeCommand struct {
	dog Dog
}

func (c *HandshakeCommand) Execute() {
	c.dog.Handshake()
}

type RunCommand struct {
	dog Dog
}

func (c *RunCommand) Execute() {
	c.dog.Run()
}

// ------- 调令者，发出命令的人
type Hoster struct {
	cmd Command
}

func (h *Hoster) Exec() {
	h.cmd.Execute()
}

func NewHoster(cmd Command) *Hoster {
	return &Hoster{
		cmd: cmd,
	}
}

func main() {
	sit := &SitCommand{dog: *NewDog("bob", "sit")}
	hand := &HandshakeCommand{dog: *NewDog("bob", "handshake")}
	run := &RunCommand{dog: *NewDog("bob", "run")}

	cmds := []Command{
		sit, hand, run,
	}

	for _, c := range cmds {
		NewHoster(c).Exec()
	}
}
