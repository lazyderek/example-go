package factory

import "testing"

func TestGenPhone(t *testing.T) {
	mobile := GenPhone("xiaomi")
	mobile.Call()
	mobile.Sms()

	mobile = GenPhone("huawei")
	mobile.Call()
	mobile.Sms()
}
