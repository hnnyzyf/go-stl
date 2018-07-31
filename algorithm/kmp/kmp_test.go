package algorithm

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestKMP(c *C) {
	k := New("AAADAAA")
	next := k.next
	res := []int{0, 1, 2, 0, 1, 2, 3}
	for i := range res {
		c.Assert(next[i], Equals, res[i])
	}

}

func (s *MySuite) TestMatch(c *C) {
	k := New("ABCDABD")
	off, ok := k.Match("BBC ABCDAB ABCDABCDABDE")
	if ok {
		c.Log("Success!")
	} else {
		c.Error("Fail!Expect 10(", off, ")")
	}
}
