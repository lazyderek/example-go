package main

import "fmt"

//建造者模式： 使用多个简单的对象一步一步构建成一个复杂的对象

type character struct {
	Name string
	Age  int
}

type builder struct {
	c character
}

func (b *builder) SetName(name string) *builder {
	b.c.Name = name
	return b
}

func (b *builder) SetAge(age int) *builder {
	b.c.Age = age
	return b
}

type director struct {
	builder builder
}

func (d *director) Create(i *info) character {
	d.builder.SetName(i.name).SetAge(i.age)
	return d.builder.c
}

func NewDirector(b builder) *director {
	return &director{
		builder: b,
	}
}

type info struct {
	name string
	age  int
}

func main() {

	builder := builder{}

	director := NewDirector(builder)

	c := director.Create(&info{name: "derek", age: 30})

	fmt.Println(c)
}
