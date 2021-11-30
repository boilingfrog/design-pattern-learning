package 解释器模式

import "testing"

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 3 + 3 + 3 + 3")
	res := p.Result().Interpret()
	expect := 13
	if res != expect {
		t.Fatalf("expect %d got %d", expect, res)
	}
	t.Log(res)
}
