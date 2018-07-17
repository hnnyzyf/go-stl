package stl

type Queue struct {
	d *Deque
}

func NewQueue() *Queue {
	return &Queue{
		d: NewDeque(),
	}
}

func (q *Queue) Push(val interface{}) {
	q.d.PushBack(val)
}

func (q *Queue) Pop() (interface{}, bool) {
	return q.d.PopFront()
}

func (q *Queue) IsEmpty() bool {
	return q.d.IsEmpty()
}

func (q *Queue) Len() int {
	return q.d.Len()
}
