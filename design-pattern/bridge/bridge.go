package main

import "fmt"

// 桥接模式： 是用于把抽象化与实现化解耦，使得二者可以独立变化。
// 这种类型的设计模式属于结构型模式，它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。
// 这种模式涉及到一个作为桥接的接口，使得实体类的功能独立于接口实现类。这两种类型的类可被结构化改变而互不影响

type Traveler interface {
	Go()
}

type Walker struct {
	tool TrafficTool
}

func (w *Walker) Go() {
	fmt.Println("去旅行...")
	w.tool.Go()
	fmt.Println("抵达目的地...")
}

type Bicycler struct {
	tool TrafficTool
}

func (b *Bicycler) Go() {
	fmt.Println("去旅行...")
	b.tool.Go()
	fmt.Println("抵达目的地...")

}

// 交通工具
type TrafficTool interface {
	Go()
}

type Footer struct {
}

func (*Footer) Go() {
	fmt.Println("徒步...")
}

type Bicycle struct {
}

func (*Bicycle) Go() {
	fmt.Println("骑自行车...")
}

func NewWalker(tool TrafficTool) *Walker {
	return &Walker{tool: tool}
}

func NewBicycler(tool TrafficTool) *Bicycler {
	return &Bicycler{tool: tool}
}

func main() {
	w := NewWalker(&Footer{})
	w.Go()

	b := NewBicycler(&Bicycle{})
	b.Go()
}
