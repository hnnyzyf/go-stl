package deque

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	d *Deque
}

var _ = Suite(&MySuite{d: New()})

func (s *MySuite) pushfront(n int) {
	for i := n - 1; i >= 0; i-- {
		s.d.PushFront(i)
	}
}

func (s *MySuite) pushback(n int) {
	for i := 0; i < n; i++ {
		s.d.PushBack(i)
	}
}

func (s *MySuite) popfront(n int) {
	for i := 0; i < n; i++ {
		s.d.PopFront()
	}
}

func (s *MySuite) popback(n int) {
	for i := 0; i < n; i++ {
		s.d.PopBack()
	}
}

func (s *MySuite) get(n int) interface{} {
	return s.d.Get(n)
}

func (s *MySuite) begin() *pos {
	return s.d.begin
}

func (s *MySuite) end() *pos {
	return s.d.end
}

func (s *MySuite) reset() {
	s.d = New()
}

func (s *MySuite) len() int {
	return s.d.Len()
}

func (s *MySuite) TestPushFront(c *C) {
	s.reset()
	n := 1025 * ChunckSize
	s.pushfront(n)
	//fmt.Println(d)
	for i := 0; i < n; i++ {
		v := s.get(i)
		c.Assert(v, Equals, i)
	}

	s.reset()
	data := []struct {
		n     int
		begin int
		end   int
	}{
		{4 * ChunckSize, 2, 5},
		{2 * ChunckSize, 1, 6},
		{1 * ChunckSize, 1, 7},
		{1 * ChunckSize, 4, 11},
		{1024 * ChunckSize, 120, 1151},
	}

	for i := range data {
		s.pushfront(data[i].n)
		c.Assert(s.begin().chunck, Equals, data[i].begin)
		c.Assert(s.end().chunck, Equals, data[i].end)
	}
}

func (s *MySuite) TestPushBack(c *C) {
	s.reset()
	n := 1025 * ChunckSize
	s.pushback(n)
	//fmt.Println(d)
	for i := 0; i < n; i++ {
		v := s.get(i)
		c.Assert(v, Equals, i)
	}

	s.reset()
	data := []struct {
		n     int
		begin int
		end   int
	}{
		{4 * ChunckSize, 2, 5},
		{2 * ChunckSize, 1, 6},
		{1 * ChunckSize, 0, 6},
		{1 * ChunckSize, 4, 11},
		{1024 * ChunckSize, 128, 1159},
	}

	for i := range data {
		s.pushback(data[i].n)
		c.Assert(s.begin().chunck, Equals, data[i].begin)
		c.Assert(s.end().chunck, Equals, data[i].end)
	}
}

func (s *MySuite) TestPop(c *C) {
	s.reset()
	n := 10 * ChunckSize
	p := 8 * ChunckSize
	s.pushback(n)
	s.popback(p)
	for i := 0; i < n-p; i++ {
		v := s.get(i)
		c.Assert(v, Equals, i)
	}

	s.reset()
	s.pushfront(n)
	s.popfront(p)
	for i := 0; i < n-p; i++ {
		v := s.get(i)
		c.Assert(v, Equals, i+p)
	}
}

func (s *MySuite) TestPushPop(c *C) {
	s.reset()
	n := 10 * ChunckSize
	p := 8 * ChunckSize
	s.pushback(n)
	s.popfront(p)
	for i := 0; i < n-p; i++ {
		v := s.get(i)
		c.Assert(v, Equals, i+p)
	}

	s.reset()
	s.pushfront(n)
	s.popback(p)
	for i := 0; i < n-p; i++ {
		v := s.get(i)
		c.Assert(v, Equals, i)
	}
}

func (s *MySuite) TestLen(c *C) {
	s.reset()

	data := []struct {
		call   func(int)
		n      int
		length int
	}{
		{s.pushback, 0, 0},
		{s.pushback, 1000 * ChunckSize, 1000 * ChunckSize},
		{s.popback, 200 * ChunckSize, 800 * ChunckSize},
		{s.pushfront, ChunckSize / 2, 800*ChunckSize + ChunckSize/2},
		{s.popback, ChunckSize / 2, 800 * ChunckSize},
		{s.popfront, 1024 * ChunckSize, 0},
		{s.popback, 1024 * ChunckSize, 0},
	}

	for i := range data {
		data[i].call(data[i].n)
		c.Assert(s.len(), Equals, data[i].length)
	}
}

func (s *MySuite) BenchmarkPushback(c *C) {
	s.reset()
	for i := 0; i < c.N; i++ {
		s.d.PushBack(i)
	}
}

func (s *MySuite) BenchmarkPushFront(c *C) {
	s.reset()
	for i := 0; i < c.N; i++ {
		s.d.PushFront(i)
	}
}
