package main

import "fmt"

// 备忘录模式： 保存一个对象的某个状态，以便在适当的时候恢复对象。备忘录模式属于行为型模式。

type Memento struct {
	state   string
	nothing string
}

type Originer struct {
	state string
}

func (o *Originer) SetState(state string) {
	o.state = state
}

func (o *Originer) Save2Memento() *Memento {
	return &Memento{
		state:   o.state,
		nothing: "123",
	}
}

type Manager struct {
	Mementos []*Memento
}

func NewManager() *Manager {
	return &Manager{
		//Mementos: make()
	}
}

func (m *Manager) Add(mm *Memento) {
	m.Mementos = append(m.Mementos, mm)
}

func (m *Manager) Get(idx int) *Memento {
	return m.Mementos[idx]
}

func (m *Manager) Remove(idx int) {
	if idx >= len(m.Mementos) {
		return
	}
	m.Mementos = append(m.Mementos[:idx-1], m.Mementos[idx:]...)
}

func main() {
	manager := NewManager()
	originer := Originer{}
	originer.SetState("996")

	// 存档
	manager.Add(originer.Save2Memento())
	fmt.Println("state", originer.state)

	// 修改状态
	originer.SetState("007")
	fmt.Println("state", originer.state)

	// 读档，恢复状态
	originer.SetState(manager.Get(0).state)
	fmt.Println("state", originer.state)
}
