package main

import (
	"container/list"
	"fmt"
)

// 观察者模式： 当对象间存在一对多关系时，则使用观察者模式（Observer Pattern）。
// 比如，当一个对象被修改时，则会自动通知它的依赖对象。观察者模式属于行为型模式。

type Subject interface {
	AddObserver(Observer)
	RmObserver(Observer)
	Notify()
}

type Observer interface {
	Happen(Subject)
}

type Subjecter struct {
	value     int
	Observers *list.List
}

func NewSubjecter() *Subjecter {
	return &Subjecter{
		value:     20000,
		Observers: list.New(),
	}
}

func (s *Subjecter) AddObserver(ob Observer) {
	s.Observers.PushBack(ob)
}

func (s *Subjecter) RmObserver(me Observer) {
	for ob := s.Observers.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(*Observer) == &me {
			s.Observers.Remove(ob)
			break
		}
	}
}

func (s *Subjecter) Notify() {
	for ob := s.Observers.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(Observer).Happen(s)
	}
}

type ObserverOne struct {
}

func (o *ObserverOne) Happen(s Subject) {
	fmt.Println("ObserverOne", s.(*Subjecter).value)
}

type ObserverTwo struct {
}

func (o *ObserverTwo) Happen(s Subject) {
	fmt.Println("ObserverTwo", s.(*Subjecter).value)
}

func main() {
	subjecter := NewSubjecter()
	one := &ObserverOne{}
	two := &ObserverTwo{}

	subjecter.AddObserver(one)
	subjecter.AddObserver(two)

	subjecter.Notify()
}
