package 桥接模式

import "testing"

func TestBridge(t *testing.T) {
	nm := NewHandsetBrandM()
	nm.SetHandsetSoft(&HandsetGame{})
	t.Log(nm.Run())
	nm.SetHandsetSoft(&HandsetAddressList{})
	t.Log(nm.Run())

	nn := NewHandsetBrandN()
	nn.SetHandsetSoft(&HandsetGame{})
	t.Log(nn.Run())
	nn.SetHandsetSoft(&HandsetAddressList{})
	t.Log(nn.Run())
}
