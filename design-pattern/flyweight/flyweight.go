package main

import (
	"fmt"
	"math/rand"
)

// 享元模式：主要用于减少创建对象的数量，以减少内存占用和提高性能。
// 这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式

type Ipoker interface {
	SetColor(color string)
	SetFace(face string)
	Show()
}

type Poker struct {
	face  string
	color string
}

func (p *Poker) Show() {
	fmt.Printf("%s%s\n", p.color, p.face)
}

func (p *Poker) SetColor(color string) {
	p.color = color
}

func (p *Poker) SetFace(face string) {
	p.face = face
}

// 享元工厂，用于存储享元对象
type FlyweightFactory struct {
	colors []string
	pokers map[string]*Poker
}

func NewFlyweight() *FlyweightFactory {
	fly := &FlyweightFactory{}
	fly.pokers = make(map[string]*Poker)
	return fly
}

var faces = []string{
	"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K",
}

func (f *FlyweightFactory) GetPoker(color string) *Poker {
	if f.pokers[color] == nil {
		fmt.Println("create color", color)
		f.pokers[color] = &Poker{color: color}
	}
	ind := rand.Intn(len(faces) - 1)
	fmt.Println(ind, rand.Intn(len(faces)-1))
	f.pokers[color].face = faces[ind]
	return f.pokers[color]
}

func main() {
	fly := NewFlyweight()
	p := fly.GetPoker("黑桃")
	p.Show()
	p = fly.GetPoker("红心")
	p.Show()
	p = fly.GetPoker("梅花")
	p.Show()
	p = fly.GetPoker("方块")
	p.Show()
}
