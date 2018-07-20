package algorithm

import (
	"testing"
)

func Test_KMP(t *testing.T) {
	k := NewKMP("ABCDABD")
	off, ok := k.Match("BBC ABCDAB ABCDABCDABDE")
	if ok {
		t.Log("Success!")
	} else {
		t.Error("Fail!Expect 10(", off, ")")
	}
}
