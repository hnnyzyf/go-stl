package stl

import (
	"fmt"
	"testing"
)

func Test_trie(t *testing.T) {
	ac := NewAhoCorasickDoubleArrayTrie()

	keywords := []string{
		"hers",
		"his",
		"she",
		"he",
	}

	ac.buildTrie(keywords)

	fmt.Println(ac.root)
}
