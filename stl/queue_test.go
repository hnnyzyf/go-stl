package stl

import (
	"fmt"
	"testing"
)

func Test_IsEmpty(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	q := NewQueue()

	for i := range a {
		q.Push(a[i])
	}

	for !q.IsEmpty() {
		x, _ := q.Pop()
		fmt.Println(x, q.d.begin, q.d.end)
	}
}
