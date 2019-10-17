package main

import "fmt"

// 模板模式： 一个抽象类公开定义了执行它的方法的方式/模板。它的子类可以按需要重写方法实现，但调用将以抽象类中定义的方式进行。
// 这种类型的设计模式属于行为型模式。

// 工作接口
type Iworker interface {
	Design()
	Working()
	Finished()
}

type Worker struct {
	Name string
	Iworker
}

// 骨架算法，定义了操作的执行方法以及流程
// Iworker的方法需要实现
func (w *Worker) StartWork() {
	fmt.Println("I'm", w.Name)
	w.Design()
	w.Working()
	w.Finished()
}

type Coder struct {
	Worker // 匿名集成Worker
}

// 实现IWorker的方法
func (c *Coder) Design() {
	fmt.Println("Design", c.Name)
}

func (c *Coder) Working() {
	fmt.Println("Working", c.Name)
}

func (c *Coder) Finished() {
	fmt.Println("Finished", c.Name)
}

// 封装一层执行接口，限制外界访问内部参数或函数
type Run interface {
	StartWork()
}

func NewCoder() Run {
	w := Coder{}
	w.Worker.Name = "Me"
	w.Worker.Iworker = &w
	return &w
}

func main() {
	w := NewCoder()
	w.StartWork()
}
