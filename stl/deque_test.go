package stl

import (
	"testing"
)

func BenchmarkPushfront(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		d.Pushfront(i)
	}
}

func BenchmarkPushback(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		d.Pushback(i)
	}
}

func BenchmarkPush(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		d.Pushfront(i)
		d.Pushback(i)
	}
}

func BenchmarkMix(b *testing.B) {
	d := NewDeque()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			d.Pushfront(i)
		} else {
			d.Pushback(i)
		}

		if i%100 == 0 && i > 100 {
			for j := 0; j < 60; j++ {
				if j%2 == 0 {
					d.Popfront()
				} else {
					d.Popback()
				}
			}
		}
	}
}
