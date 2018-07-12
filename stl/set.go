package stl

//Set is a RBtree
type Set struct {
	r *RBtree
}

func NewSet() *Set {
	return &Set{
		r: NewRBTree(),
	}
}

//Insert add a new Value into Set if not exist or do nothing if exist
func (s *Set) Insert(val Value) {
	s.r.Push(val)
}

//Erase lete a Value if exist
func (s *Set) Erase(val Value) {
	s.r.Pop(val)
}

//Iterator will return a Iterator
func (s *Set) Iterator(isAsc bool) func() (Value, bool) {
	return s.r.Iterator(isAsc)
}

//Len return the size of rbtree
func (s *Set) Len() int {
	return s.r.Len()
}

//we provide some Val to use

type uint64Val uint64

func (i intVal) Less(v Value) bool {
	return i < v.(intVal)
}

func (i intVal) More(v Value) bool {
	return i > v.(intVal)
}

type stringVal string

func (i stringVal) Less(v Value) bool {
	return i < v.(stringVal)
}

func (i stringVal) More(v Value) bool {
	return i > v.(stringVal)
}

type float64Val float64

func (i float64Val) Less(v Value) bool {
	return i < v.(float64Val)
}

func (i float64Val) More(v Value) bool {
	return i > v.(float64Val)
}

type RuneVal rune

func (i RuneVal) Less(v Value) bool {
	return i < v.(RuneVal)
}

func (i RuneVal) More(v Value) bool {
	return i > v.(RuneVal)
}
