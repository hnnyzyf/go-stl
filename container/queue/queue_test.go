package queue

import (
	"testing"
)

func Test_IsEmpty(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	q := New()

	for i := range a {
		q.Push(a[i])
	}

	for !q.IsEmpty() {
		q.Pop()
	}
}
