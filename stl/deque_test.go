package stl

import (
	"testing"
)

func Test_Front(t *testing.T) {
	d := NewDeque()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i := range data {
		d.PushFront(data[i])
	}

	for i := range res {
		if v, ok := d.PopFront(); v != res[i] || !ok {
			t.Error("Except:", res[i], "(", v, ")")
		}
	}
}

func Test_Back(t *testing.T) {
	d := NewDeque()
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i := range data {
		d.PushBack(data[i])
	}

	for i := range res {
		if v, ok := d.PopBack(); v != res[i] || !ok {
			t.Error("Except:", res[i], "(", v, ")")
		}
	}
}

func Test_Len(t *testing.T) {
	d := NewDeque()
	for i := 0; i < 100; i++ {
		d.PushBack(i)
	}

	if d.Len() != 100 {
		t.Error("Fail!Expect 100(", d.Len(), ")")
	}

	for i := 0; i < 50; i++ {
		d.PopFront()
	}

	if d.Len() != 50 {
		t.Error("Fail!Expect 50(", d.Len(), ")")
	}

	for i := 0; i < 50; i++ {
		d.PopBack()
	}

	if d.Len() != 0 {
		t.Error("Fail!Expect 0(", d.Len(), ")")
	}

	for i := 0; i < 50; i++ {
		d.PopBack()
	}

	if d.Len() != 0 {
		t.Error("Fail!Expect 0(", d.Len(), ")")
	}

}

func BenchmarkPushFront(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
	}
}

func BenchmarkPushBack(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
}

func BenchmarkPush(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
		d.PushBack(i)
	}
}

func BenchmarkMix(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			d.PushFront(i)
		} else {
			d.PushBack(i)
		}

		if i%100 == 0 && i > 100 {
			for j := 0; j < 60; j++ {
				if j%2 == 0 {
					d.PopFront()
				} else {
					d.PopBack()
				}
			}
		}
	}
}
