package treemap

import (
	"github.com/hnnyzyf/go-stl/container/pair"
	"testing"
)

type kv struct {
	k int
	v int
}

func Test_Insert(t *testing.T) {
	treemap := New()
	a := []kv{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
		{2, 1},
	}
	res := []kv{
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 2},
	}

	for i := range a {
		treemap.Insert(pair.Int(a[i].k, a[i].v))
	}

	j := 0
	for i := treemap.Begin(); i.LessEqual(treemap.End()); i.Next() {
		p := i.GetValue()
		k := p.GetKey().(int)
		v := p.GetValue().(int)
		if k != res[j].k || v != res[j].v {
			t.Error("Fail,Expect ", res[j].k, ":", res[j].v, "(", k, ":", v, ")")
		}
		j++
	}

}

func Test_Erase(t *testing.T) {
	treemap := New()
	a := []kv{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
		{2, 1},
	}

	del := []kv{
		{3, 0},
		{3, 0},
		{4, 1},
	}

	res := []struct {
		k int
		v int
	}{
		{1, 0},
		{2, 1},
	}

	for i := range a {
		treemap.Insert(pair.Int(a[i].k, a[i].v))
	}

	for i := range del {
		treemap.Erase(pair.Int(del[i].k, del[i].v))
	}

	j := 0
	for i := treemap.Begin(); i.LessEqual(treemap.End()); i.Next() {
		p := i.GetValue()
		k := p.GetKey().(int)
		v := p.GetValue().(int)
		if k != res[j].k || v != res[j].v {
			t.Error("Fail,Expect ", res[j].k, ":", res[j].v, "(", k, ":", v, ")")
		}
		j++
	}

}

func Test_Find(t *testing.T) {
	treemap := New()
	a := []kv{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
		{2, 1},
	}

	res := []kv{
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 2},
	}

	for i := range a {
		treemap.Insert(pair.Int(a[i].k, a[i].v))
	}

	for i := range res {
		p := treemap.Find(pair.Int(res[i].k, res[i].v))

		pk := p.GetKey().(int)
		py := p.GetValue().(int)

		rk := res[i].k
		ry := res[i].v

		if pk != rk || py != ry {
			t.Error("Fail,Expect ", res[i], "(", p, ")")
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	t := New()
	for i := 0; i < b.N; i++ {
		e := pair.Int(i, i)
		t.Insert(e)
	}
}

func BenchmarkFind(b *testing.B) {
	t := New()
	for i := 0; i < 100000; i++ {
		t.Insert(pair.Int(i, i))
	}

	for i := 0; i < b.N; i++ {
		t.Find(pair.Int(i, i))
	}
}
