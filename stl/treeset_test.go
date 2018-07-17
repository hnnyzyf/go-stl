package stl

import (
	"fmt"
	"testing"
)

func Test_Iterator(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := NewTreeSet()

	for i := range a {
		s.Insert(IntVal(a[i]))
	}

	next := s.Iterator(true)

	for e, ok := next(); ok; e, ok = next() {
		fmt.Println(e)
	}
}
