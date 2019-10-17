package prototype

import "fmt"

//原型模式 用于重复创建一个对象

type Dog struct {
	Name   string
	Gender string
	Color  string
}

func (d *Dog) Run() {
	fmt.Println("The dog is running")
}

func NewDog() *Dog {
	return &Dog{}
}
