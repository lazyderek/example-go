package factory

import (
	"Design-Patterns-go/factory/huawei"
	"Design-Patterns-go/factory/xiaomi"
)

type Phone interface {
	Call()
	Sms()
}

func GetXiaomi() Phone {
	return &xiaomi.XiaoMi{}
}

func GetHuaWei() Phone {
	return &huawei.HuaWei{}
}

func GenPhone(phoneName string) Phone {
	switch phoneName {
	case "xiaomi":
		return GetXiaomi()
	case "huawei":
		return GetHuaWei()
	default:
		panic("Invalid phone Name")
	}
}
