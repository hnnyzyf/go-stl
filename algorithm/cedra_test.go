package algorithm

import (
	"fmt"
	"testing"
)

type info struct {
	sibling rune
	child   rune
}

func Test_t(t *testing.T) {
	a := &info{11, 12}
	b := &info{22, 23}
	r := &a.sibling

	fmt.Println(r, *r)

	r = &b.child
	fmt.Println(r, *r, b)
	*r = 25
	fmt.Println(r, *r, b)
}
