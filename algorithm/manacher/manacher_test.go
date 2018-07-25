package algorithm

import (
	"testing"
)

func Test_manacher(t *testing.T) {
	text := "12212921"
	m := NewManacher('$', '#')
	//fmt.Println(text)
	x := m.GetLPS(text)
	if x != "12921" {
		t.Error("Fail!Expect 12921(", len(x), ")")
	}
}
