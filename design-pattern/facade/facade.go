package main

import "fmt"

// 外观模式： 隐藏系统的复杂性，并向客户端提供了一个客户端可以访问系统的接口。
// 这种类型的设计模式属于结构型模式，它向现有的系统添加一个接口，来隐藏系统的复杂性。
// 这种模式涉及到一个单一的类，该类提供了客户端请求的简化方法和对现有系统类方法的委托调用。
// 关键： 在客户端和复杂系统之间再加一层，这一层将调用顺序、依赖关系等处理好。 client -> facade -> system

type SystemA struct {
}

func (a *SystemA) Access() {
	fmt.Println("access A")
}

type SystemB struct {
}

func (a *SystemB) Access() {
	fmt.Println("access B")
}

type SystemC struct {
}

func (a *SystemC) Access() {
	fmt.Println("access C")
}

type Facade struct {
	A SystemA
	B SystemB
	C SystemC
}

func (f *Facade) AccessA() {
	f.A.Access()
}

func (f *Facade) AccessB() {
	f.B.Access()
}

func (f *Facade) AccessC() {
	f.C.Access()
}

func main() {
	f := new(Facade)

	f.AccessA()
	f.AccessB()
	f.AccessC()
}
