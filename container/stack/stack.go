package stack

import (
	"github.com/hnnyzyf/go-stl/container/deque"
)

type Stack struct {
	d *deque.Deque
}

func New() *Stack {
	return &Stack{
		d: deque.New(),
	}
}

//Push insert a element
func (s *Stack) Push(val interface{}) {
	s.d.PushBack(val)
}

//Pop remove a element
func (s *Stack) Pop() (interface{}, bool) {
	return s.d.PopBack()
}

//Top return next element
func (s *Stack) Top() (interface{}, bool) {
	return s.d.Last()
}

//IsEpty test whether container is empty
func (s *Stack) IsEmpty() bool {
	return s.d.IsEmpty()
}

//Len return size
func (s *Stack) Len() int {
	return s.d.Len()
}
