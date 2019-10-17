package huawei

import "fmt"

type HuaWei struct {
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

func (*HuaWei) Call() {
	fmt.Println("HuaWei call ....")
}

func (*HuaWei) Sms() {
	fmt.Println("HuaWei sms ....")
}
