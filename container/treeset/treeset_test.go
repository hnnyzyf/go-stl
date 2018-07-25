package treeset

import (
	"github.com/hnnyzyf/go-stl/container/value"
	"testing"
)

func Test_Iterator(t *testing.T) {
	a := []int{9, 2, 3, 4, 5, 1, 7, 8, 6}
	res := []value.Int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set := New()

	for i := range a {
		set.Insert(value.Int((a[i])))
	}

	j := 0
	for i := set.Begin(); i.LessEqual(set.End()); i.Next() {
		v := i.GetValue()
		if v.(value.Int) != res[j] {
			t.Error("Fail,Expect ", res[j], "(", v, ")")
		}
		j++
	}
}
