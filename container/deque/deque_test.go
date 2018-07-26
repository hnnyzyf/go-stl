package deque

import (
	//"fmt"
	"testing"
)

func Test_Front(t *testing.T) {
	d := New()
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
	d := New()
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

func pushback(d *Deque, n int) {
	for i := 0; i < n; i++ {
		d.PushBack(n)
	}
}

func pushfront(d *Deque, n int) {
	for i := 0; i < n; i++ {
		d.PushFront(n)
	}
}

func popback(d *Deque, n int) {
	for i := 0; i < n; i++ {
		d.PopBack()
	}
}

func popfront(d *Deque, n int) {
	for i := 0; i < n; i++ {
		d.PopFront()
	}
}

func Test_Len(t *testing.T) {
	d := New()

	data := []struct {
		call func(*Deque, int)
		n    int
		len  int
	}{
		{pushback, 1024 * ChunckSize, 1024 * ChunckSize},
		{popfront, 500 * ChunckSize, 524 * ChunckSize},
		{pushfront, 101 * ChunckSize, 625 * ChunckSize},
		{popfront, 500 * ChunckSize, 125 * ChunckSize},
		{popback, 100 * ChunckSize, 25 * ChunckSize},
	}

	for i := range data {
		data[i].call(d, data[i].n)
		if d.Len() != data[i].len {
			t.Error("Fail!Expect ", data[i].len, "(", d.Len(), ")")
		}
	}
}

func Test_reallocmmap(t *testing.T) {
	d := New()
	data := []struct {
		i     int
		begin int
		end   int
	}{
		{4 * ChunckSize, 4, 7},
		{1, 2, 6},
		{2*ChunckSize - 1, 2, 7},
		{1, 1, 7},
		{ChunckSize - 1, 1, 7},
		{1, 0, 7},
		{ChunckSize - 1, 0, 7},
		{1, 4, 12},
		{4*ChunckSize - 1, 4, 15},
		{1, 2, 14},
		{2*ChunckSize - 1, 2, 15},
		{1, 1, 15},
		{ChunckSize - 1, 1, 15},
		{1, 0, 15},
		{ChunckSize - 1, 0, 15},
		{1008 * ChunckSize, 0, 1023},
		{1, 128, 1152},
	}

	for i := range data {
		pushback(d, data[i].i)
		//fmt.Println(d.begin, d.end, d.Len())
		if d.begin.chunck != data[i].begin || d.end.chunck != data[i].end {
			t.Error(i, ":Fail！Except ", data[i].begin, ":", data[i].end, "(", d.begin.chunck, ":", d.end.chunck, ")")
		}
	}

	data = []struct {
		i     int
		begin int
		end   int
	}{
		{1, 128, 1151},
		{385*ChunckSize - 1, 128, 767},
		{1, 160, 798},
		//{1, 1, 643},
	}

	for i := range data {
		popback(d, data[i].i)
		//fmt.Println(d.begin, d.end, len(d.mmap), d.Len())
		if d.begin.chunck != data[i].begin || d.end.chunck != data[i].end {
			t.Error(i, ":Fail！Except ", data[i].begin, ":", data[i].end, "(", d.begin.chunck, ":", d.end.chunck, ")")
		}
	}

}

func BenchmarkPushFront(b *testing.B) {
	d := New()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
	}
}

func BenchmarkPushBack(b *testing.B) {
	d := New()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
}

func BenchmarkPush(b *testing.B) {
	d := New()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
		d.PushBack(i)
	}
}

func BenchmarkMix(b *testing.B) {
	d := New()
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
