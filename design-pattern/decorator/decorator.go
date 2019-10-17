package main

import (
	"fmt"
	"time"
)

// 装饰器模式： 允许向一个现有的对象添加新的功能，同时又不改变其结构。
// 这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

// 这里的timer是一个装饰器，对someFunc进行装饰，计算someFunc的运行时长
func timer(someFunc func()) func() {
	f := func() {
		fmt.Println("this is a decorator")
		start := time.Now()
		someFunc()
		end := time.Now()
		fmt.Println("time: ", end.Sub(start))
	}
	return f
}

// 被装饰者， 懵逼的me函数
func Me() {
	fmt.Println("it is me")
}

func main() {
	me := timer(Me)
	me()
}
