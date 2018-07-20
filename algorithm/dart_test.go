package algorithm

import (
	//"fmt"
	"testing"
)

type Hit struct {
	//the beginning index, inclusive.
	begin int
	//the ending index, exclusive
	end int
	//the value assigned to the keyword
	value rune
}

func NewHit(begin int, end int, value rune) *Hit {
	return &Hit{
		begin: begin,
		end:   end,
		value: value,
	}
}

func Test_trie(t *testing.T) {
	ac := NewDart()

	keywords := []string{
		"space",
		"keyword",
		"ch",
	}

	ac.Build(keywords)

	//fmt.Println(ac.base[134])
	//fmt.Println(len(ac.check))

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
	}

	for i := range data {
		if ac.Matches(data[i].text) != data[i].res {
			t.Error("Fail(", data[i].text, "),Expect:", data[i].res)
		}
	}

}
