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

func Test_IsRBtree(t *testing.T) {
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

func Test_AscIter(t *testing.T) {
	rb := NewRBTree()

	a := []intVal{12, 1, 9, 2, 11, 7, 0, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}
	res := []intVal{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for _, x := range a {
		rb.Push(x)
	}

	e := rb.End()
	i := 0
	for b := rb.Begin(); !b.Equal(e); b.Next() {
		v := b.Value().(intVal)
		if v != res[i] {
			t.Error("Fail,Except", res[i], "(", v, ")")
		}
		i++
	}
}

func Test_DescIter(t *testing.T) {
	rb := NewRBTree()

	a := []intVal{12, 1, 9, 2, 0, 11, 7, 19, 4, 15, 18, 5, 14, 13, 10, 16, 6, 3, 8, 17}
	res := []intVal{19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for _, x := range a {
		rb.Push(x)
	}

	b := rb.Begin()
	i := 0
	for e := rb.End(); !e.Equal(b); e.Last() {
		v := e.Value().(intVal)
		if v != res[i] {
			t.Error("Fail,Except", res[i], "(", v, ")")
		}
		i++
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

func BenchmarkGet(b *testing.B) {
	rb := NewRBTree()

	for i := 0; i < 1000000; i++ {
		rb.Push((intVal)(i))
	}

	for i := 0; i < b.N; i++ {
		_, _ = rb.Get((intVal)(i))
	}

}

func BenchmarkIterator(b *testing.B) {
	rb := NewRBTree()
	for i := 0; i < 10000000; i++ {
		rb.Push((intVal)(i))
	}

	b.N = 10000000
	iter := rb.Begin()

	for i := 0; i < b.N; i++ {
		iter.Next()
	}

}
