package stl

type Stack struct {
	d *Deque
}

func NewStack() *Stack {
	return &Stack{
		d: NewDeque(),
	}
}

func (s *Stack) Push(val interface{}) {
	s.d.Pushback(val)
}

func (s *Stack) Pop() (interface{}, bool) {
	return s.d.PopBack()
}
