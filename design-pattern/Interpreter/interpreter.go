package main

import (
	"fmt"
)

// 解释器模式： 提供了评估语言的语法或表达式的方式，它属于行为型模式。
// 这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等

func main() {
	//unicode.IsDigit("1")
	fmt.Println(int(1 - '0'))
}
