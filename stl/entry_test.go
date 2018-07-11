package stl

import (
	"testing"
)

func Test_boolEntry(t *testing.T) {
	tree := NewTreeMap()
	a := []*BoolEntry{
		{true, 0},
		{true, 32},
		{false, "222"},
		{true, "111111111"},
	}
	res := []*BoolEntry{
		{true, "111111111"},
		{false, "222"},
	}

	for i := range a {
		tree.Put(a[i])
	}

	next := tree.Iterator(true)

	for i := range res {
		n, ok := next()
		k := n.GetKey().(bool)
		v := n.GetValue().(string)
		if k != res[i].k || v != res[i].v || !ok {
			t.Error("Fail,Expect ", res[i].k, ":", res[i].v, "(", k, ":", v, ")")
		}
	}
}
