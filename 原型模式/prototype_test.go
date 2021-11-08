package 原型模式

import "testing"

var (
	deepCopyManager    *PrototypeManager
	shallowCopyManager *PrototypeManager
)

// 深拷贝实现Cloneable
type DeepCopy struct {
	name string
}

func (t *DeepCopy) Clone() Cloneable {
	tc := *t
	return &tc
}

// 浅拷贝实现Cloneable
type ShallowCopy struct {
	name string
}

func (t *ShallowCopy) Clone() Cloneable {
	return t
}

func TestDeepCopyClone(t *testing.T) {
	t1 := deepCopyManager.Get("dc")

	t2 := t1.Clone()
	// 深拷贝，指向的不是同一个变量的地址
	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}

	t21 := t2.(*DeepCopy)
	t21.name = "ShallowCopy-test"

	t11 := t1.(*DeepCopy)
	// 深拷贝name，不会影响到copy前的变量
	if t11.name == t21.name {
		t.Fatal("shallowCopy err")
	}
}

func TestShallowCopyClone(t *testing.T) {
	t1 := shallowCopyManager.Get("sc")

	t2 := t1.Clone()
	// 浅拷贝，变量地址的指向不变
	if t1 != t2 {
		t.Fatal("error! get clone not working")
	}

	t21 := t2.(*ShallowCopy)
	t21.name = "ShallowCopy-test"

	t11 := t1.(*ShallowCopy)
	// 深拷贝name，copy之前的变量和copy之后的变量同时更改
	if t11.name != t21.name {
		t.Fatal("shallowCopy err")
	}
}

func init() {
	deepCopyManager = NewPrototypeManager()

	dc := &DeepCopy{
		name: "deepCopy",
	}
	deepCopyManager.Set("dc", dc)

	shallowCopyManager = NewPrototypeManager()
	sc := &ShallowCopy{
		name: "shallowCopy",
	}
	shallowCopyManager.Set("sc", sc)
}
