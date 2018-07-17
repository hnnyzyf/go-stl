package stl

import (
	"testing"
)

func Test_trie(t *testing.T) {
	ac := NewDart()

	keywords := []string{
		"space",
		"keyword",
		"ch",
	}

	ac.Build(keywords)

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
		if ac.Matchs(data[i].text) != data[i].res {
			t.Error("Fail,Expect:", data[i].res)
		}
	}

}
