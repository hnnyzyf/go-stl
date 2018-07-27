package algorithm

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/hnnyzyf/go-stl/container/queue"
	"github.com/hnnyzyf/go-stl/container/value"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	ac *Dart

	text io.RuneReader

	result []int
}

var _ = Suite(NewMysuit())

func NewMysuit() *MySuite {
	s := &MySuite{}
	s.reset()
	return s
}

func (s *MySuite) reset() {
	s.ac = NewDart()
}

func (s *MySuite) match(text string) bool {
	return s.ac.Matches(text)
}

func (s *MySuite) build(keywords []string) error {
	return s.ac.Build(keywords)
}

func (s *MySuite) parseString(text string) {
	s.result = make([]int, 0)
	s.ac.ParseString(text, s.result1)
}

func (s *MySuite) parseText() {
	s.ac.ParseText(s.text, nil)
}
func (s *MySuite) result1(output []int) {
	s.result = append(s.result, output...)
}

func (s *MySuite) result2(output []int) {
	for _, i := range output {
		fmt.Println("Find", s.ac.keywords[i])
	}
}

func (s *MySuite) getstate() map[int]rune {
	res := make(map[int]rune)
	q := queue.New()
	q.Push(s.ac.root)
	for !q.IsEmpty() {
		e, _ := q.Pop()
		s := e.(*state)
		res[s.id] = s.c
		fmt.Println("getstate", s.id, "(", s.c, ")")
		for i := s.success.Begin(); i.LessEqual(s.success.End()); i.Next() {
			n := i.GetData().GetValue().(*state)
			q.Push(n)
		}
	}

	return res
}

func (s *MySuite) getemit() map[int][]int {
	res := make(map[int][]int)
	q := queue.New()
	q.Push(s.ac.root)
	for !q.IsEmpty() {
		e, _ := q.Pop()
		s := e.(*state)
		res[s.id] = make([]int, 0)
		for i := s.emits.Begin(); i.LessEqual(s.emits.End()); i.Next() {
			v := int(i.GetData().(value.Int))
			res[s.id] = append(res[s.id], v)
		}
		fmt.Println("getemit", s.id, "(", res[s.id], ")")
		for i := s.success.Begin(); i.LessEqual(s.success.End()); i.Next() {
			n := i.GetData().GetValue().(*state)
			q.Push(n)
		}
	}

	return res
}

func (s *MySuite) getcheck() map[int]int {
	res := make(map[int]int)
	for i, v := range s.ac.check {
		if v != 0 {
			res[i] = v
			fmt.Println("getcheck", i, "(", v, ")")
		}
	}

	return res
}

func (s *MySuite) getbase() map[int]int {
	res := make(map[int]int)
	for i, v := range s.ac.base {
		if v != 0 {
			res[i] = v
			fmt.Println("getbase", i, "(", v, ")")
		}
	}

	return res
}

func (s *MySuite) loadCN() error {
	s.reset()
	keywords, err := loadDictionary("./resource/cn/dictionary.txt")
	if err != nil {
		return err
	}
	s.build(keywords)

	text, err := loadText("./resource/cn/text.txt")
	if err != nil {
		return err
	}

	s.text = text
	return nil
}

func (s *MySuite) loadEN() error {
	s.reset()
	keywords, err := loadDictionary("./resource/en/dictionary.txt")
	if err != nil {
		return err
	}
	s.build(keywords)
	text, err := loadText("./resource/en/text.txt")
	if err != nil {
		return err
	}
	s.text = text
	return nil
}

//loadDictionary read key from a text file
func loadDictionary(path string) ([]string, error) {
	var (
		file   *os.File
		err    error
		part   []byte
		prefix bool
	)
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	lines := make([]string, 0)
	text := bufio.NewReader(file)
	buffer := bytes.NewBuffer([]byte{})

	for {
		part, prefix, err = text.ReadLine()
		if err != nil {
			break
		}

		if _, err = buffer.Write(part); err != nil {
			return nil, err
		}

		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}

	if err == io.EOF {
		return lines, nil
	} else {
		return nil, err
	}
}

func loadText(path string) (io.RuneReader, error) {
	var (
		file *os.File
		err  error
	)

	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(file), nil

}

func (s *MySuite) TestMatches(c *C) {
	s.reset()

	keywords := []string{
		"space",
		"keyword",
		"ch",
		"ab",
		"abc",
		"cde",
		"cdef",
	}

	if err := s.build(keywords); err != nil {
		c.Error(err)
	}

	data := []struct {
		text string
		res  bool
	}{
		{"space", true},
		{"keyword", true},
		{"ch", true},
		{"  ch", true},
		{"chkeyword", true},
		{"oooospace2", true},
		{"c", false},
		{"", false},
		{"spac", false},
		{"nothing", false},
		{"abc", true},
		{"abcdefghijkabcdef", true},
	}

	for i := range data {
		c.Assert(s.match(data[i].text), Equals, data[i].res)
	}

}

func (s *MySuite) TestBuild(c *C) {
	s.reset()
	keywords, err := loadDictionary("./resource/cn/dictionary.txt")
	if err != nil {
		c.Error(err)
	}
	err = s.build(keywords)
	if err != nil {
		c.Error(err)
	}
}

func (s *MySuite) TestParseString(c *C) {
	s.reset()

	keywords := []string{
		"space",
		"keyword",
		"ch",
		"ab",
		"abc",
		"cde",
		"cdef",
	}

	if err := s.build(keywords); err != nil {
		c.Error(err)
	}

	txt := "abcdefghijkabcdef"

	s.parseString(txt)

	res := []int{3, 4, 5, 6, 3, 4, 5, 6}

	for i := range res {
		c.Assert(res[i], Equals, s.result[i])
	}

}

func (s *MySuite) TestParseText(c *C) {
	if err := s.loadCN(); err != nil {
		c.Error(err)
	}

	s.parseText()
}
