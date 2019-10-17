package xiaomi

import "fmt"

type XiaoMi struct {
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

func (*XiaoMi) Call() {
	fmt.Println("XiaoMi call ....")
}

func (*XiaoMi) Sms() {
	fmt.Println("XiaoMi sms ....")
}
