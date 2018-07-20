package stl

//TreeSet is a RBtree
type TreeSet struct {
	r *RBTree
}

func NewTreeSet() *TreeSet {
	return &TreeSet{
		r: NewRBTree(),
	}
}

//Insert add a new Value into TreeSet if not exist or do nothing if exist
func (s *TreeSet) Insert(val Value) {
	s.r.Push(val)
}

//Erase lete a Value if exist
func (s *TreeSet) Erase(val Value) {
	s.r.Pop(val)
}

//Iterator will return a Iterator
func (s *TreeSet) Begin() *Riterator {
	return s.r.Begin()
}

func (s *TreeSet) End() *Riterator {
	return s.r.End()
}

//Len return the size of rbtree
func (s *TreeSet) Len() int {
	return s.r.Len()
}

//we provide some Val to use

type Uint64Val uint64

func (i Uint64Val) Less(v Value) bool {
	return i < v.(Uint64Val)
}

func (i Uint64Val) More(v Value) bool {
	return i > v.(Uint64Val)
}

type IntVal uint64

func (i IntVal) Less(v Value) bool {
	return i < v.(IntVal)
}

func (i IntVal) More(v Value) bool {
	return i > v.(IntVal)
}

type StringVal string

func (i StringVal) Less(v Value) bool {
	return i < v.(StringVal)
}

func (i StringVal) More(v Value) bool {
	return i > v.(StringVal)
}

type Float64Val float64

func (i Float64Val) Less(v Value) bool {
	return i < v.(Float64Val)
}

func (i Float64Val) More(v Value) bool {
	return i > v.(Float64Val)
}

type RuneVal rune

func (i RuneVal) Less(v Value) bool {
	return i < v.(RuneVal)
}

func (i RuneVal) More(v Value) bool {
	return i > v.(RuneVal)
}
