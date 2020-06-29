package dp

import "testing"

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}

// type1 实现了Cloneable接口
// 那么，我们可以克隆type1了
func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatalf("Get clone not work")
	}
}
