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
