package stl

import (
	"sort"
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

//state is a node of a trie tree
type state struct {

	//goto table
	success map[rune]*state
	//output table
	emits map[int]bool
	//if Fail,turn to this state
	failure *state

	//the max emit in emits
	semit []int
	//the depth of current state,also the length of words
	depth int
	//the index of keywords
	index int
}

func newState() *state {
	return &state{
		success: make(map[rune]*state),
		emits:   make(map[int]bool),

		semit: make([]int, 0),
	}
}

//addEmit add a new pattern
func (s *state) addEmit(keyword int) {
	if _, ok := s.emits[keyword]; !ok {
		s.emits[keyword] = true
		s.semit = append(s.semit, keyword)
		//do a sort
		sort.Ints(s.semit)
	}
}

//addEmits add more than one new patter
func (s *state) addEmits(keywords []int) {
	for i := range keywords {
		s.addEmit(keywords[i])
	}
}

//isAcceptable return whether the state is the final state
func (s *state) isAcceptable() bool {
	return s.depth > 0 && len(s.emits) != 0
}

//nextState return next state
//if root node call thie function,ignoreRootState is true,else false
func (s *state) nextState(c rune, ignoreRootState bool) *state {
	next, ok := s.success[c]
	if !ignoreRootState && !ok && s.depth == 0 {
		return s
	}
	return next
}

//addState will add a state
func (s *state) addState(c rune) *state {
	next, ok := s.success[c]
	if !ok {
		n := newState()
		n.depth = s.depth + 1
		s.success[c] = n
		return n
	}

	return next
}

func (s *state) getMaxEmit() int {
	if !sort.IntsAreSorted(s.semit) {
		sort.Ints(s.semit)
	}

	if len(s.semit) > 0 {
		return s.semit(len(s.semit) - 1)
	} else {
		return 0
	}
}

type AhoCorasickDoubleArrayTrie struct {

	//base array of the Double Array Trie structure
	base []int
	//check array of the Double Array Trie structure
	check []int
	//fail table of the Aho Corasick automata
	fail []int
	//output table of the Aho Corasick automata
	output [][]int
	//outer value array
	v []rune
	//the length of every key
	l []int
	//the size of base and check array
	size int

	//the root state of trie
	root *state
	//whether the position has been used
	used []bool
	//the size of the key-pair sets
	keySize int
}

func NewAhoCorasickDoubleArrayTrie() *AhoCorasickDoubleArrayTrie {
	return &AhoCorasickDoubleArrayTrie{
		check:  make([]int, 0),
		base:   make([]int, 0),
		fail:   make([]int, 0),
		output: make([][]int, 0),
		v:      make([]rune, 0),
		l:      make([]int, 0),

		root: newState(),
		used: make([]bool, 0),
	}
}

//Build will create a ACDAT
func (a *AhoCorasickDoubleArrayTrie) Build(keys []string) {
	//create a basic tries
	a.buildTrie(keys)

	//create a double array tries
	a.buildDAT(len(keys))

	//create a fail table
	//a.buildAC()

	//clean
	//a.used = nil
	//a.root = nil
}

//buildTrie will create a basic tries state tree
func (a *AhoCorasickDoubleArrayTrie) buildTrie(keys []string) {
	for i, key := range keys {
		a.addKey(key, i)
	}
}

//buildTrie will create a basic tries state tree
func (a *AhoCorasickDoubleArrayTrie) addKey(key string, index int) {
	b := []rune(key)

	//create a trie tree
	curr := a.root
	for i := range b {
		curr = curr.addState(b[i])
	}

	//add the keyword index
	curr.addEmit(index)

	//record the length of keywords
	a.l = append(a.l, len(b))
}

//buildDAT will create a double array tries
func (a *AhoCorasickDoubleArrayTrie) buildDAT(size int) {
	//init
	a.keySize = size
	a.addbase(0, 1)
	a.addcheck(0, 0)

	//get all children
	slibings := a.fetch(a.root)

	//add state
	a.insert(slibings)

}

//addBase will add a new base state
func (a *AhoCorasickDoubleArrayTrie) addbase(idx int, val int) {
	if len(a.base) <= idx {
		b := make([]int, idx-len(a.base)+1)
		a.base = append(a.base, b...)
	}

	a.base[idx] = val
}

//addCheck will add a new check
func (a *AhoCorasickDoubleArrayTrie) addcheck(idx int, val int) {
	if len(a.check) <= idx {
		b := make([]int, idx-len(a.check)+1)
		a.check = append(a.check, b...)
	}

	a.check[idx] = val
}

//fetch return all children
func (a *AhoCorasickDoubleArrayTrie) fetch(s *state) map[rune]*state {
	slibings := make(map[rune]*state)

	//the root state or final state
	if s.isAcceptable() {
		//create a fake state as the final state
		fake := newState()
		fake.depth == -s.depth-1
		fake.addEmit(s.getMaxEmit())
		slibings['0'] = fake
	}

	//add other state
	for key, _ := range s.success {
		slibings[key] = s.success[key]
	}

	return slibings
}

//insert add some state into base and check table
func (a *AhoCorasickDoubleArrayTrie) insert(slibings map[rune]*state) {
	//each state has a different base index
	var begin int = 0

}
