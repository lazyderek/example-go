package main

import "fmt"

// 抽象工厂，相当于一个工厂的集合。一个工厂可以生产多个产品，如手机、电脑等，因为每个产品都是一个接口，所以一个工厂会可以有多个接口。

const (
	PorductA = "A"
	PorductB = "B"
)

// --------- 产品A接口 ---------
type ProductMobile interface {
	CreateMobile()
}

// --------- 产品B接口 ---------
type ProductComputer interface {
	CreateComputer()
}

// 实现你的product
// ----------- mobile -------------
type ProductFactoryA struct {
	BrandName string `json:"brand_name"`
}

func (m *ProductFactoryA) CreateMobile() {
	fmt.Println("generate a mobile : " + m.BrandName)
}
func (c *ProductFactoryA) CreateComputer() {
	fmt.Println("generate a computer : " + c.BrandName)
}

// ----------- mobile -------------
type ProductFactoryB struct {
	BrandName string `json:"brand_name"`
}

func (m *ProductFactoryB) CreateMobile() {
	fmt.Println("generate a mobile : " + m.BrandName)
}

func (c *ProductFactoryB) CreateComputer() {
	fmt.Println("generate a computer : " + c.BrandName)
}

// --------- 抽象工厂 ---------
type IAbstractFactory interface {
	CreateMobile()
	CreateComputer()
}

// 创建一个工厂
func NewAbstractFactory(product string) IAbstractFactory {
	switch product {
	case PorductA:
		return &ProductFactoryA{BrandName: PorductA}
	case PorductB:
		return &ProductFactoryB{BrandName: PorductB}
	}
	return nil
}

func main() {
	a := NewAbstractFactory(PorductA)
	a.CreateComputer()
	a.CreateMobile()

	b := NewAbstractFactory(PorductB)
	b.CreateComputer()
	b.CreateMobile()
}
