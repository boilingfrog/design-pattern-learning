package 桥接模式

import "testing"

func TestBridge(t *testing.T) {
	hw := NewHuawei()
	hw.SetHandsetSoft(&HandsetGame{})
	t.Log(hw.Run())
	hw.SetHandsetSoft(&HandsetAddressList{})
	t.Log(hw.Run())

	ap := NewApple()
	ap.SetHandsetSoft(&HandsetGame{})
	t.Log(ap.Run())
	ap.SetHandsetSoft(&HandsetAddressList{})
	t.Log(ap.Run())
}
