package stl

import (
	"errors"
	"fmt"
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
	success *TreeMap
	//output table
	emits *TreeSet
	//if Fail,turn to this state
	failure *state

	//the depth of current state,also the length of words
	depth int
	//the index of keywords
	index int
}

func newState() *state {
	return &state{
		success: NewTreeMap(),
		emits:   NewTreeSet(),
	}
}

//addEmit add a new pattern
func (s *state) addEmit(keyword int) {
	s.emits.Insert((IntVal)(keyword))
}

//addEmit add a new pattern
func (s *state) addEmits(emits *TreeSet) {
	next := emits.Iterator(true)

	for e, ok := next(); ok; e, ok = next() {
		s.emits.Insert(e.(IntVal))
	}
}

//isAcceptable return whether the state is the final state
func (s *state) isFinalState() bool {
	return s.depth > 0 && s.emits.Len() != 0
}

//nextState return next state
//if root node call thie function is true,else false
func (s *state) nextState(c rune) *state {
	//get next
	next, _ := s.success.Get(NewRuneEntry(c, 0))
	return next.(*state)
}

//addState will add a state
func (s *state) addState(c rune) *state {
	next, ok := s.success.Get(NewRuneEntry(c, 0))
	if !ok {
		n := newState()
		n.depth = s.depth + 1
		s.success.Put(NewRuneEntry(c, n))
		return n
	}

	return next.(*state)
}

func (s *state) getMaxEmit() int {
	if s.emits.Len() == 0 {
		panic("Impossible State")
	}

	//get a iterator
	next := s.emits.Iterator(false)

	//get the max emit if exist
	if max, ok := next(); ok {
		return (int)(max.(IntVal))
	} else {
		panic("impossible state")
	}
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
	err = da.buildAC()

	if err != nil {
		return err
	}

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
		da.addKey(key, i)
	}

	return nil
}

//addKey add a new state node
func (da *Dart) addKey(key string, index int) {
	b := []rune(key)

	if len(b) == 0 {
		panic("Failed to insert zero-length key")
	}

	//create a trie tree
	curr := da.root
	for i := range b {
		curr = curr.addState(b[i])
	}

	//add the keyword index
	curr.addEmit(index)

	//record the length of keywords
	da.l = append(da.l, len(b))

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

//buildDAT create a double array trie
func (da *Dart) buildDAT(size int) error {
	//alloc 8MB memory
	da.resize(65536 * 32)
	da.base[0] = 1
	da.check[0] = 0
	da.keySize = size

	slibings := da.fetch(da.root)

	da.insert(slibings)

	return nil
}

//fetch return all children state
func (da *Dart) fetch(s *state) []*RuneEntry {
	slibings := make([]*RuneEntry, 0)

	//add other
	next := s.success.Iterator(false)

	for e, ok := next(); ok; e, ok = next() {
		slibings = append(slibings, e.(*RuneEntry))
	}

	if s.isFinalState() {
		slibings = append(slibings, NewRuneEntry('0', s.fake()))
	}

	return slibings
}

//insert retrun the begin value
func (da *Dart) insert(slibings []*RuneEntry) int {
	begin := 0
	pos := int(slibings[0].Key()) + 2
	if da.nextCheckPos > pos {
		pos = da.nextCheckPos - 1
	} else {
		pos = pos - 1
	}

	nonzero_num := 0
	first := 0

	if da.allocSize <= pos {
		da.resize(pos + 1)
	}

	// 此循环体的目标是找出满足base[begin + a1...an]  == 0的n个空闲空间,a1...an是siblings中的n个节点
	for {
		pos++

		if da.allocSize <= pos {
			da.resize(pos + 1)
		}

		if da.check[pos] != 0 {
			nonzero_num++
			continue
		} else if first == 0 {
			da.nextCheckPos = pos
			first = 1
		} else {
			//dothing
		}

		begin := pos - int(slibings[0].Key())

		if da.allocSize <= begin+int(slibings[len(slibings)-1].Key()) {
			l := float64(da.keySize) / float64(da.progress+1)
			if l < 1.05 {
				l = 1.05
			}
			da.resize(int(float64(da.allocSize) * l))
		}

		if da.used[begin] {
			continue
		}

		for i := range slibings {
			if da.check[begin+int(slibings[i].Key())] != 0 {
				continue
			}
		}

		break
	}
	// -- Simple heuristics --
	// if the percentage of non-empty contents in check between the
	// index
	// 'next_check_pos' and 'check' is greater than some constant value
	// (e.g. 0.9),
	// new 'next_check_pos' index is written by 'check'.

	if float64(nonzero_num)/float64(pos-da.nextCheckPos+1) >= 0.95 {
		da.nextCheckPos = pos
	}

	da.used[begin] = true

	if da.size < begin+int(slibings[len(slibings)-1].Key())+1 {
		da.size = begin + int(slibings[len(slibings)-1].Key()) + 1
	}

	for i := range slibings {
		da.check[begin+int(slibings[i].Key())] = begin
	}

	for i := range slibings {
		s := slibings[i].GetValue().(*state)
		new_sliblings := da.fetch(s)
		if len(new_sliblings) == 0 {
			da.base[begin+int(slibings[i].Key())] = -s.getMaxEmit() - 1
			da.progress++
		} else {
			h := da.insert(new_sliblings)
			da.base[begin+int(slibings[i].Key())] = h
		}

		s.index = begin + int(slibings[i].Key())
	}

	return begin

}

//buildAC create a fail table
func (da *Dart) buildAC() error {
	da.fail = make([]int, da.size+1)
	da.fail[1] = da.base[0]
	da.output = make([][]int, da.size+1)

	queue := NewQueue()

	// 第一步，将深度为1的节点的failure设为根节点
	next := da.root.success.Iterator(true)

	for e, ok := next(); ok; e, ok = next() {
		s := e.GetValue().(*state)
		s.setFailure(da.root, da.fail)
		queue.Push(s)
		da.addOutput(s)
	}

	// 第二步，为深度 > 1 的节点建立failure表，这是一个bfs
	for {
		n, ok := queue.Pop()
		if !ok {
			break
		}
		fmt.Println("curr", n)
		curr := n.(*state)
		next := curr.success.Iterator(true)
		for e, ok := next(); ok; e, ok = next() {
			transition := e.(*RuneEntry).Key()
			target := e.GetValue().(*state)
			if target == nil {
				panic("target is nil")
			}
			queue.Push(target)

			fail := curr.failure
			for fail.nextState(transition) == nil {
				fail = fail.failure
			}

			fmt.Println("fail2", fail.nextState(transition))
			nfail := fail.nextState(transition)
			target.setFailure(nfail, da.fail)
			target.addEmits(nfail.emits)
			da.addOutput(target)

		}

	}

	return nil
}

func (da *Dart) addOutput(s *state) {
	output := make([]int, 0)
	next := s.emits.Iterator(true)

	for {
		if e, ok := next(); ok {
			output = append(output, int(e.(IntVal)))
		} else {
			break
		}
	}

	da.output[s.index] = output
}

func (da *Dart) Matchs(text string) bool {
	b := []rune(text)

	curr := 0

	for i := range b {
		curr = da.getState(curr, b[i])
		hit := da.output[curr]
		if hit != nil {
			return true
		}
	}
	return false
}

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
	p := b + int(c) + 1

	if b != da.check[p] {
		if pos == 0 {
			return 0
		}
		return -1
	}

	return p
}
