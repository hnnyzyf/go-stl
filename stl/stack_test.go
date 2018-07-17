package stl

import (
	"testing"
)

func BenchmarkPushMix(b *testing.B) {
	s := NewStack()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkMix(b *testing.B) {
	s := NewStack()

	for i := 0; i < b.N; i++ {
		s.Push(i)
		s.Pop()
	}
}

func Test_IsEmpty(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Error("Fail!Expect true(false)")
	}

	if s.Len() != 0 {
		t.Error("Fail!Expect 0(", s.Len(), ")")
	}

	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	if s.Len() != 1000 {
		t.Error("Fail!Expect 1000(", s.Len(), ")")
	}

	for i := 0; i < 1000; i++ {
		s.Pop()
	}

	if !s.IsEmpty() {
		t.Error("Fail!Expect true(false)")
	}

	if s.Len() != 0 {
		t.Error("Fail!Expect 0(", s.Len(), ")")
	}

}
