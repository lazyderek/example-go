package singleton

import "sync"

type User struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Age      uint   `json:"age"`
	Phone    string `json:"phone"`
}

/*
单例模式： 确保某对象只被实例化一次，该对象是公用的
*/

// singleUser只被实例化一次
var singleUser *User
var once sync.Once

// 非线程的情况
func GetUser() *User {
	if singleUser == nil {
		singleUser = &User{Username: "123"}
	}
	return singleUser
}

// 存在线程的情况
func GetSingleUser() *User {
	once.Do(func() {
		singleUser = &User{Username: "345"}
	})
	return singleUser
}
