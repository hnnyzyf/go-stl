package queue

import (
	"github.com/hnnyzyf/go-stl/container/deque"
)

type Queue struct {
	d *deque.Deque
}

func New() *Queue {
	return &Queue{
		d: deque.New(),
	}
}

//Push insert element
func (q *Queue) Push(val interface{}) {
	q.d.PushBack(val)
}

//Pop remove next element
func (q *Queue) Pop() (interface{}, bool) {
	return q.d.PopFront()
}

//IsEmpty test whether container is empty
func (q *Queue) IsEmpty() bool {
	return q.d.IsEmpty()
}

//Len return size
func (q *Queue) Len() int {
	return q.d.Len()
}

//Back access last element
func (q *Queue) Back() interface{} {
	if q.IsEmpty() {
		return nil
	} else {
		return q.d.Get(q.Len() - 1)
	}
}

//Front access next element
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	} else {
		return q.d.Get(0)
	}
}
