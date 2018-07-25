package treemap

import (
	"github.com/hnnyzyf/go-stl/container/pair"
	"github.com/hnnyzyf/go-stl/container/rbtree"
	"github.com/hnnyzyf/go-stl/container/value"
)

//TreeMap is a RBTree
type TreeMap struct {
	r *rbtree.RBTree
}

func New() *TreeMap {
	return &TreeMap{
		r: rbtree.New(),
	}
}

//Insert will add a key into TreeMap if not exist or update a key if exists
func (t *TreeMap) Insert(e pair.Pair) {
	t.r.Push(e.(value.Value))
}

//Erase will delete a key if exists or do nothing
func (t *TreeMap) Erase(e pair.Pair) {
	t.r.Pop(e.(value.Value))
}

//Find will return value if exists a key or nil if  not exist
//it is a replace of []
func (t *TreeMap) Find(e pair.Pair) pair.Pair {
	if n := t.r.Find(e.(value.Value)); n != nil {
		return n.(pair.Pair)
	} else {
		return nil
	}
}

//Get is a
func (t *TreeMap) Get(i *Miterator) pair.Pair {
	return i.iter.GetValue().(pair.Pair)

}

//Begin return iterator to beginning
func (t *TreeMap) Begin() *Miterator {
	return newMiterator(t.r.Begin())
}

//End return iterator to end
func (t *TreeMap) End() *Miterator {
	return newMiterator(t.r.End())
}

//Len return the size of rbtree
func (t *TreeMap) Len() int {
	return t.r.Len()
}

//Empty test whether container is empty
func (t *TreeMap) IsEmpty() bool {
	return t.r.Len() == 0
}

//Miterator is the iterator of treemap
type Miterator struct {
	iter *rbtree.Riterator
}

func newMiterator(i *rbtree.Riterator) *Miterator {
	return &Miterator{
		iter: i,
	}
}

//GetValue is a replace of *
func (i *Miterator) GetValue() pair.Pair {
	return i.iter.GetValue().(pair.Pair)
}

//Next is a replace of ++
func (i *Miterator) Next() {
	i.iter.Next()
}

//last is a replace of --
func (i *Miterator) Last() {
	i.iter.Last()
}

//Equal is a replace of ==
func (i *Miterator) Equal(m *Miterator) bool {
	return i.iter.Equal(m.iter)
}

//LessEqual is a replace of <=
func (i *Miterator) LessEqual(m *Miterator) bool {
	return i.iter.LessEqual(m.iter)
}

//MoreEqual is a replace of >=
func (i *Miterator) MoreEqual(m *Miterator) bool {
	return i.iter.MoreEqual(m.iter)
}
