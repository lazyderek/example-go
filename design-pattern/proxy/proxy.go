package main

import "fmt"

// 代理模式： 一个类代表另一个类的功能。这种类型的设计模式属于结构型模式

type IHttp interface {
	Post()
	Get()
}

// 被代理者
type http struct {
}

// 被代理者的功能
func (*http) Post() {
	fmt.Println("http post ...")
}

func (*http) Get() {
	fmt.Println("http get ...")
}

func (*http) Put() {
	fmt.Println("http put ...")
}

// 代理者
type ProxyHttp struct {
	http *http
}

func NewHttpProxy(t *http) IHttp {
	return &ProxyHttp{
		http: t,
	}
}

// 将被代理者的功能再封一层，添加其他的操作
func (p *ProxyHttp) Post() {
	fmt.Println("proxy do something before post ...")
	p.http.Post()
	fmt.Println("proxy do something after post ...")
}

func (p *ProxyHttp) Get() {
	fmt.Println("proxy do something before Get ...")
	p.http.Get()
	fmt.Println("proxy do something after Get ...")
}

// 被代理者Put()方法并没有被代理实现，这样可以限制外界http的put请求

func main() {
	proxy := NewHttpProxy(&http{})

	proxy.Get()

	proxy.Post()
}
