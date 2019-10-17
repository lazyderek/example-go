package main

import "fmt"

const (
	OrderValue   = 1
	CookingValue = 2
	SendingValue = 3
	FinishValue  = 4
)

type Context struct {
	StateValue   int
	CurrentState IState
}

func NewContext() *Context {
	return &Context{
		StateValue: OrderValue,
	}
}

func (c *Context) SetStateValue(value int) {
	c.StateValue = value
	c.changeState()
}

func (c *Context) changeState() {
	switch c.StateValue {
	case OrderValue:
		c.CurrentState = &StateOrder{}
	case CookingValue:
		c.CurrentState = &StateCooking{}
	case SendingValue:
		c.CurrentState = &StateSending{}
	case FinishValue:
		c.CurrentState = &StateFinished{}
	default:
		c.CurrentState = &StateOrder{}
	}
	c.CurrentState.Do()
}

type IState interface {
	Do()
}

type StateOrder struct {
}

func (*StateOrder) Do() {
	fmt.Println("下单中...")
}

type StateCooking struct {
}

func (*StateCooking) Do() {
	fmt.Println("正在做菜...")
}

type StateSending struct {
}

func (*StateSending) Do() {
	fmt.Println("外卖配送中...")
}

type StateFinished struct {
}

func (*StateFinished) Do() {
	fmt.Println("外卖已收到...")
}

func main() {
	c := NewContext()

	c.SetStateValue(OrderValue)
	c.SetStateValue(CookingValue)
	c.SetStateValue(SendingValue)
	c.SetStateValue(FinishValue)
}
