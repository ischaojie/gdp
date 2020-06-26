package solid

import "testing"

func TestDIP(t *testing.T) {

	zhangSan := Driver{}
	benz := Benz{}
	bmw := BMW{}
	zhangSan.Drive(benz) // 奔驰汽车开始运行
	zhangSan.Drive(bmw)  // 宝马汽车开始运行
}
