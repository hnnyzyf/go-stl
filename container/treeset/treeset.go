package treeset

import (
	"github.com/hnnyzyf/go-stl/container/rbtree"
	"github.com/hnnyzyf/go-stl/container/value"
)

//TreeSet is a RBtree
type TreeSet struct {
	r *rbtree.RBTree
}

func New() *TreeSet {
	return &TreeSet{
		r: rbtree.New(),
	}
}

//Insert add a new Value into TreeSet if not exist or do nothing if exist
func (s *TreeSet) Insert(val value.Value) {
	s.r.Push(val)
}

//Erase lete a Value if exist
func (s *TreeSet) Erase(val value.Value) {
	s.r.Pop(val)
}

//Iterator will return a Iterator
func (s *TreeSet) Begin() *rbtree.Riterator {
	return s.r.Begin()
}

func (s *TreeSet) End() *rbtree.Riterator {
	return s.r.End()
}

//Len return the size of rbtree
func (s *TreeSet) Len() int {
	return s.r.Len()
}

//Find will return value if exists a key or nil if  not exist
func (s *TreeSet) Find(val value.Value) bool {
	if v := s.r.Find(val); v != nil {
		return true
	} else {
		return false
	}

}
