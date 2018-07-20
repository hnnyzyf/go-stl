package stl

import (
	"testing"
)

func Test_Iterator(t *testing.T) {
	a := []int{9, 2, 3, 4, 5, 1, 7, 8, 6}
	res := []IntVal{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := NewTreeSet()

	for i := range a {
		tree.Insert(IntVal(a[i]))
	}

	e := tree.End()
	i := 0
	for b := tree.Begin(); !b.Equal(e); b.Next() {
		v := b.Value().(IntVal)
		if v != res[i] {
			t.Error("Fail,Expect ", res[i], "(", v, ")")
		}
		i++
	}
}
