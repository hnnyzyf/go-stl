package algorithm

import (
	"errors"
	"github.com/hnnyzyf/go-stl/stl"
	"fmt"
)

//state is a node of a trie tree
type state struct {
	//the id in base and check table
	id int
	//the depth of current state,also the length of keyword
	depth int
	//goto table
	success *stl.TreeMap
	//if Fail,turn to this state
	failure *state

	//output table,record the keyword index in input texts
	emits *stl.TreeSet
}

func newState() *state {
	return &state{
		success: stl.NewTreeMap(),
		emits:   stl.NewTreeSet(),
	}
}

//addEmit add a new keyword
func (s *state) addEmit(keyword int) {
	s.emits.Insert((stl.IntVal)(keyword))
}

//addEmits add many pattern
func (s *state) addEmits(emits *stl.TreeSet) {
	for b := emits.Begin(); b != emits.End(); b.Next() {
		s.emits.Insert(b.Value())
	}

}

//isFinalState return whether the state is the final state
func (s *state) isFinalState() bool {
	return s.depth > 0 && s.emits.Len() != 0
}

//nextState return next state
func (s *state) nextState(c rune) *state {
	//get next
	next, ok := s.success.Get(stl.NewRuneEntry(c, 0))
	if ok {
		//find the nextstate
		return next.(*state)
	} else if !ok && s.depth == 0 {
		//don't find and current state is root
		return s
	} else {
		//return nil if nothing find
		return nil
	}
}

//addState will add a state
func (s *state) addState(c rune) *state {
	next, ok := s.success.Get(stl.NewRuneEntry(c, 0))
	if !ok {
		n = newState()
		n.depth = s.depth + 1
		s.success.Put(stl.NewRuneEntry(c, n))
		return n
	}

	return next.(*state)
}

//return the maximum index
func (s *state) getMaxEmit() int {
	if s.emits.Len() == 0 {
		panic("Impossible State")
	}

	//get the Max Emit
	e := s.emits.End()
	return int(e.(IntVal))
}

func (s *state) fake() *state {
	//create a fake state,which depth<0 and success is null as the final state
	f := newState()
	f.depth = -s.depth - 1
	f.addEmit(s.getMaxEmit())
	return f
}

func (s *state) setFailure(f *state, fail []int) {
	s.failure = f
	fail[s.index] = f.index
}

type Dart struct {

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
	//a parameter controls the memory growth speed of the dynamic array
	progress int
	//the next position to check unused memory
	nextCheckPos int
	//the allocSize of the dynamic array
	allocSize int
	//the size of the key-pair sets
	keySize int
}

func NewDart() *Dart {
	return &Dart{
		check: make([]int, 0),
		base:  make([]int, 0),
		fail:  make([]int, 0),

		output: make([][]int, 0),
		v:      make([]rune, 0),
		l:      make([]int, 0),
		root:   newState(),
		used:   make([]bool, 0),
	}
}

//Build will create a ACDAT
//the input keys should be ordered
func (da *Dart) Build(keys []string) error {
	//create a basic tries
	err := da.buildTrie(keys)
	if err != nil {
		return err
	}

	//create a double array tries
	err = da.buildDAT(len(keys))
	if err != nil {
		return err
	}

	//create a fail table
	//err = da.buildAC()

	//if err != nil {
		//return err
	//}

	//clean
	//da.used = nil
	//da.root = nil
	return nil
}

//buildTrie will create a basic tries state tree
func (da *Dart) buildTrie(keys []string) error {
	if len(keys) == 0 {
		return errors.New("No key to insert!")
	}

	for i, key := range keys {
		if err := da.addKey(key, i); err != nil {
			return err
		}
	}

	return nil
}

//addKey add a new state node
func (da *Dart) addKey(key string, index int) error {
	b := []rune(key)

	if len(b) == 0 {
		return errors.New("Failed to insert zero-length key")
	}

	//create a trie tree
	curr := da.root
	for i := range b {
		curr = curr.addState(b[i])
	}

	//add the keyword index
	curr.addEmit(index)

	return nil

}


//buildDAT create a double array trie
func (da *Dart) buildDAT(size int) error {
	//alloc 8MB memory
	da.resize(65536*32)
	//init
	da.base[0] = 1
	da.check[0] = 0

	stack:=stl.NewStack()

	return nil
}

//resize will alloc memory for base,check,used array
func (da *Dart) resize(newSize int) {
	newbase := make([]int, newSize)
	newcheck := make([]int, newSize)
	newused := make([]bool, newSize)

	if da.allocSize > 0 {
		copy(newbase[:len(da.base)], da.base)
		copy(newcheck[:len(da.check)], da.check)
		copy(newused[:len(da.used)], da.used)
	}

	da.base = newbase
	da.check = newcheck
	da.used = newused

	da.allocSize = newSize
}

//buildAC create a fail table
func (da *Dart) buildAC() error {
	da.fail = make([]int, da.size+1)
	da.fail[1] = da.base[0]
	da.output = make([][]int, da.size+1)
	queue := stl.NewQueue()

	//step 1: set the failure of state whose depth is 1 as the root state
	next := da.root.success.Iterator(true)
	for e, ok := next(); ok; e, ok = next() {
		//set the root state as failure
		s := e.GetValue().(*state)
		s.setFailure(da.root, da.fail)

		//add into queue
		queue.Push(s)

		da.addOutput(s)
	}

	//step 2，create failure state for all states whose depth > 1
	for !queue.IsEmpty() {
		//get current node
		n, _ := queue.Pop()
		curr := n.(*state)

		//bfs
		next := curr.success.Iterator(true)
		for e, ok := next(); ok; e, ok = next() {
			//the char we meet
			c := e.(*stl.RuneEntry).Key()

			//find the next fail state which have the same char in success table
			fail := curr.failure
			for fail.nextState(c) == nil {
				fail = fail.failure
			}
			fail = fail.nextState(c)

			//add child into queue
			child := e.GetValue().(*state)
			queue.Push(child)

			//set failure
			child.setFailure(fail, da.fail)
			child.addEmits(fail.emits)
			da.addOutput(child)
		}

	}

	return nil
}

func (da *Dart) addOutput(s *state) {
	output := make([]int, 0)
	next := s.emits.Iterator(true)

	for e, ok := next(); ok; e, ok = next() {
		output = append(output, int(e.(stl.IntVal)))
	}

	da.output[s.index] = output
}

func (da *Dart) Matches(text string) bool {
	b := []rune(text)

	curr := 0

	for i := range b {
		curr = da.getState(curr, b[i])
		hit := da.output[curr]
		if len(hit) > 0 {
			return true
		}
	}
	return false
}

//get state
func (da *Dart) getState(curr int, c rune) int {
	ncurr := da.transitionWithRoot(curr, c)
	for ncurr == -1 {
		curr = da.fail[curr]
		ncurr = da.transitionWithRoot(curr, c)
	}

	return ncurr
}

func (da *Dart) transitionWithRoot(pos int, c rune) int {
	b := da.base[pos]
	p := b + int(c)

	v := 0
	if p < len(da.check) {
		v = da.check[p]
	}
	if b != v {
		if pos == 0 {
			return 0
		}
		return -1
	}

	return p
}
