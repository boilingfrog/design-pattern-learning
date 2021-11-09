package 代理模式

import "testing"

func TestPursuitGirl(t *testing.T) {
	pr := NewProxy("小红")
	t.Log(pr.GiveFollowers())
	t.Log(pr.GiveDolls())
	t.Log(pr.GiveChocolate())
}
