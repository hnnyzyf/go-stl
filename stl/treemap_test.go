package stl

import (
	"testing"
)

type kv struct {
	k int
	v int
}

func new(k int, v int) *kv {
	return &kv{k, v}
}

func (k *kv) Less(v Value) bool {
	return k.k < v.(*kv).k
}

func (k *kv) Equal(v Value) bool {
	return k.k == v.(*kv).k
}

func (k *kv) More(v Value) bool {
	return k.k > v.(*kv).k
}

func (k *kv) GetKey() interface{} {
	return k.k
}

func (k *kv) GetValue() interface{} {
	return k.v
}

func Test_Put(t *testing.T) {
	tree := NewTreeMap()
	a := []*kv{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
		{2, 1},
	}
	res := []*kv{
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 2},
	}

	for i := range a {
		tree.Put(a[i])
	}

	next := tree.Iterator(true)

	for i := range res {
		n, ok := next()
		k := n.GetKey().(int)
		v := n.GetValue().(int)
		if k != res[i].k || v != res[i].v || !ok {
			t.Error("Fail,Expect ", res[i].k, ":", res[i].v, "(", k, ":", v, ")")
		}
	}

}

func Test_Delete(t *testing.T) {
	tree := NewTreeMap()
	a := []*kv{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
		{2, 1},
	}

	del := []*kv{
		{3, 0},
		{3, 0},
		{4, 1},
	}

	res := []*kv{
		{1, 0},
		{2, 1},
	}

	for i := range a {
		tree.Put(a[i])
	}

	for i := range del {
		tree.Delete(del[i])
	}

	next := tree.Iterator(true)

	for i := range res {
		n, ok := next()
		k := n.GetKey().(int)
		v := n.GetValue().(int)
		if k != res[i].k || v != res[i].v || !ok {
			t.Error("Fail,Expect ", res[i].k, ":", res[i].v, "(", k, ":", v, ")")
		}
	}

}

func Test_Get(t *testing.T) {
	tree := NewTreeMap()
	a := []*kv{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{4, 1},
		{4, 2},
		{2, 1},
	}

	res := []*kv{
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 2},
	}

	for i := range a {
		tree.Put(a[i])
	}

	for i := range res {
		v, ok := tree.Get(res[i])

		if !ok || v.(int) != res[i].v {
			t.Error("Fail,Expect ", res[i].v, "(", v, ")")
		}
	}
}

func BenchmarkPut(b *testing.B) {
	t := NewTreeMap()
	for i := 0; i < b.N; i++ {
		e := new(i, i)
		t.Put(e)
	}
}

func BenchmarkGet(b *testing.B) {
	t := NewTreeMap()
	for i := 0; i < 1000000; i++ {
		e := new(i, i)
		t.Put(e)
	}

	for i := 0; i < b.N; i++ {
		e := new(i, i)
		t.Get(e)
	}
}
