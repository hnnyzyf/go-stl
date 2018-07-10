package stl

import (
	"testing"
)

type intVal int

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

//test root
func testRoot(n *node) bool {
	if n.color == RED {
		return false
	}
	return true
}

//test Rednode
func testRedNode(n *node) bool {
	if n.color == BLACK {
		return true
	}

	if isBlack(n.l) && isBlack(n.r) {
		return true
	} else {
		return false
	}
}

func calPath(n *node) (int, bool) {
	if n == nil {
		return 0, true
	}

	if n.color == BLACK {
		lc, ok1 := calPath(n.l)
		rc, ok2 := calPath(n.r)
		if lc == rc && ok1 && ok2 {
			return lc + 1, true
		} else {
			return -1, false
		}
	} else {
		lc, ok1 := calPath(n.l)
		rc, ok2 := calPath(n.r)
		if lc == rc && ok1 && ok2 {
			return lc, true
		} else {
			return -1, false
		}
	}
}

func testPath(n *node) bool {
	_, ok := calPath(n)
	return ok
}

func TestIsRBtree(t *testing.T) {
	rb := NewRBTree()

	a := []intVal{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}

	for _, x := range a {
		rb.Push(x)
		if ok, err := rb.IsRBTree(testRoot, testRedNode, testPath); !ok {
			t.Error(err)
			break
		}
	}

	for _, x := range a {
		rb.Pop(x)
		if ok, err := rb.IsRBTree(testRoot, testRedNode, testPath); !ok {
			t.Error(x, err)
		}
	}

	b := []stringVal{"12", " 1", " 9", " 2", " 0", " 11", " 7", " 19", " 4", " 15", " 18", " 5", " 14", " 13", " 10", " 16", " 6", " 3", " 8", " 17"}
	for _, x := range b {
		rb.Push(x)
		if ok, err := rb.IsRBTree(testRoot, testRedNode, testPath); !ok {
			t.Error(err)
		}
	}

	for _, x := range b {
		rb.Pop(x)
		if ok, err := rb.IsRBTree(testRoot, testRedNode, testPath); !ok {
			t.Error(err)
		}
	}
}

func BenchmarkPush(b *testing.B) {
	rb := NewRBTree()

	for i := 0; i < b.N; i++ {
		rb.Push((intVal)(i))
	}

}

func BenchmarkMix(b *testing.B) {
	rb := NewRBTree()

	for i := 0; i < b.N; i++ {
		rb.Push((intVal)(i))
	}

	for i := 0; i < b.N; i++ {
		rb.Pop((intVal)(i))
	}

}
