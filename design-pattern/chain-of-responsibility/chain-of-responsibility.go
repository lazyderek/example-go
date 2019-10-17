package main

import "fmt"

// 责任链模式：责任链模式（Chain of Responsibility Pattern）为请求创建了一个接收者对象的链。
// 这种模式给予请求的类型，对请求的发送者和接收者进行解耦。这种类型的设计模式属于行为型模式。

const (
	LevelInfo  = 1
	LevelDebug = 2
	LevelError = 3
)

type ILogger interface {
	Print(level uint, msg string)
	SetNextLogger(l ILogger)
}

type Logger struct {
	level uint
	next  ILogger
}

func NewLogger(level uint) ILogger {
	return &Logger{level: level}
}

func (i *Logger) Print(level uint, msg string) {
	if i.level < level && i.next != nil {
		i.next.Print(level, msg)
		return
	}
	fmt.Println(i.level, msg)
}

func (i *Logger) SetNextLogger(l ILogger) {
	i.next = l
}

func main() {
	log1 := NewLogger(LevelInfo)
	log2 := NewLogger(LevelDebug)
	log3 := NewLogger(LevelError)

	log1.SetNextLogger(log2)
	log2.SetNextLogger(log3)

	log1.Print(LevelDebug, "打印log")
}
