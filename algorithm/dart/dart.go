package algorithm

import (
	"errors"
	"io"

	"github.com/hnnyzyf/go-stl/container/pair"
	"github.com/hnnyzyf/go-stl/container/queue"
	"github.com/hnnyzyf/go-stl/container/stack"
	"github.com/hnnyzyf/go-stl/container/treemap"
	"github.com/hnnyzyf/go-stl/container/treeset"
	"github.com/hnnyzyf/go-stl/container/value"
)

//state is a node of a trie tree
type state struct {
	c rune
	//the id in base and check table
	id int
	//the depth of current state,also the length of keyword
	depth int
	//goto table
	success *treemap.TreeMap
	//if Fail,turn to this state
	failure *state

	//output table,record the keyword index in input texts
	emits *treeset.TreeSet
	//fake node
	fake *state
}

func newState() *state {
	return &state{
		success: treemap.New(),
		emits:   treeset.New(),
	}
}

//addEmit add a new keyword
func (s *state) addEmit(keyword int) {
	s.emits.Insert(value.Int(keyword))
}

//addEmits add many pattern
func (s *state) addEmits(emits *treeset.TreeSet) error {
	for i := emits.Begin(); i.LessEqual(emits.End()); i.Next() {
		v := i.GetData()
		if v == nil {
			return errors.New("An non-empty treeset return nil data")
		}
		s.emits.Insert(v)
	}
	return nil
}

//isFinalState return whether the state is the final state
func (s *state) isFinalState() bool {
	return s.depth > 0 && s.emits.Len() != 0
}

//isFakeState return whether the state is a fake state
func (s *state) isFakeState() bool {
	return s.depth < 0
}

//addFakeState create a fake state,which depth<0 and success is null as the final state
func (s *state) addFakeState() *state {
	f := newState()
	f.depth = -s.depth - 1
	f.addEmit(s.getMaxEmit())
	s.fake = f
	return f
}

//nextState return next state
func (s *state) nextState(c rune) (*state, error) {
	//get next
	p := s.success.Find(pair.Rune(c, 0))
	if p == nil && s.depth == 0 {
		return s, nil
	} else if p == nil && s.depth != 0 {
		return nil, nil
	} else if p != nil {
		next, ok := p.GetValue().(*state)
		if !ok {
			return nil, errors.New("The next state is a nil state")
		}
		return next, nil
	} else {
		panic("Impossible state")
	}
}

//addState will add a state
func (s *state) addState(c rune) (*state, error) {
	if p := s.success.Find(pair.Rune(c, 0)); p != nil {
		next, ok := p.GetValue().(*state)
		if !ok {
			return nil, errors.New("The next state we find is a nil state")
		}
		return next, nil
	} else {
		n := newState()
		n.c = c
		n.depth = s.depth + 1
		s.success.Insert(pair.Rune(c, n))
		return n, nil
	}
}

//return the maximum index
func (s *state) getMaxEmit() int {
	//get the Max Emit
	i := s.emits.End()
	if v, ok := i.GetData().(value.Int); ok {
		return int(v)
	} else {
		return 0
	}
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

	//the root state of trie
	root *state
	//whether the position has been used
	used []bool
	//the allocSize of the dynamic array
	allocSize int
	//nextPos recorad last begin
	nextPos int
	//record keywords
	keywords []string
}

func NewDart() *Dart {
	return &Dart{
		check: make([]int, 0),
		base:  make([]int, 0),
		fail:  make([]int, 0),

		output: make([][]int, 0),
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
	err = da.buildDAT()
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
	da.keywords = keys
	return nil
}

//ParseText check a text file,when hit a pattern,use callback

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
		if next, err := curr.addState(b[i]); err != nil {
			return err
		} else {
			curr = next
		}
	}

	//add the keyword index
	curr.addEmit(index)

	return nil

}

//buildDAT create a double array trie
func (da *Dart) buildDAT() error {
	//alloc 8MB memory
	da.resize(1)

	//we use double stack to do dfs
	s := stack.New()
	output := stack.New()

	s.Push(da.root)
	//step1:calculate check table in pre order
	for !s.IsEmpty() {
		n, ok := s.Pop()
		if !ok {
			return errors.New("An non-emptry stack returns nil")
		}
		curr, ok := n.(*state)
		if !ok {
			return errors.New("Find next state is a nil state")
		}
		output.Push(curr)
		//get all child
		slibings, err := da.fetch(curr)
		if err != nil {
			return err
		}

		//find a position for
		begin := da.calculate(slibings)

		//set check table and stack
		for i := range slibings {
			slibling := slibings[i]
			pos := int(slibling.GetKey().(rune)) + begin
			da.check[pos] = begin
			//set stack
			if temp, ok := slibling.GetValue().(*state); ok {
				temp.id = pos
				s.Push(temp)
			} else {
				return errors.New("Find child state is a nil state")
			}
		}
	}

	//calculate base table in post order
	//bacause if we want to calculate a base value of a state
	//we need calculate the begin value of all its children at first
	for !output.IsEmpty() {
		n, ok := output.Pop()
		if !ok {
			return errors.New("An non-emptry stack returns nil")
		}
		s, ok := n.(*state)
		if !ok {
			return errors.New("It seems we insert a nil state?")
		}
		if s.isFakeState() {
			da.base[s.id] = -s.getMaxEmit() - 1
		} else if s.isFinalState() {
			h := s.fake.id
			da.base[s.id] = h
		} else {
			n := s.success.Get(s.success.Begin())
			if n == nil {
				return errors.New("The first child of A non-empty map is nil")
			}
			if child, ok := n.GetValue().(*state); ok {
				h := child.id - int(child.c)
				da.base[s.id] = h
			} else {
				return errors.New("The first child of A non-empty map is not nil,but it store a nil state")
			}
		}

	}

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

//fecth will return slibings
func (da *Dart) fetch(s *state) ([]*pair.RunePair, error) {
	slibings := make([]*pair.RunePair, 0)

	if s.isFinalState() {
		slibings = append(slibings, pair.Rune(0, s.addFakeState()))
	}

	//add
	for i := s.success.Begin(); i.LessEqual(s.success.End()); i.Next() {
		p, ok := i.GetData().(*pair.RunePair)
		if !ok {
			return nil, errors.New("Fetch an none empty map but return nil")
		}
		slibings = append(slibings, p)
	}

	return slibings, nil
}

//calculate calculate the begin
func (da *Dart) calculate(slibings []*pair.RunePair) int {
	begin := da.nextPos
	maxKey := 0
	if len(slibings) > 0 {
		maxKey = int(slibings[len(slibings)-1].GetKey().(rune))
	}

	//find the begin
	for {
		begin++

		//alloc enough memory
		pos := maxKey + begin
		if pos >= len(da.check) {
			da.resize(pos + pos/4)
		}

		if da.used[begin] {
			continue
		}
		//find a position which check[begin+a1....an]=0
		flag := 0
		for i := range slibings {
			slibing := slibings[i]
			pos := int(slibing.GetKey().(rune)) + begin
			if da.check[pos] != 0 {
				flag = 1
				break
			}
		}

		//check begin whether it has not used
		if flag == 0 {
			da.used[begin] = true
			break
		}

	}

	da.nextPos = begin
	return begin
}

//buildAC create a fail table
func (da *Dart) buildAC() error {
	//create fail table and output table
	da.fail = make([]int, len(da.base))
	da.output = make([][]int, len(da.base))

	q := queue.New()
	//step 1:set the failure of all children of root to root
	for i := da.root.success.Begin(); i.LessEqual(da.root.success.End()); i.Next() {
		n := i.GetData()
		if n == nil {
			return errors.New("An none empty state has children but return nil")
		}
		temp, ok := n.GetValue().(*state)
		if !ok {
			return errors.New("It seems we insert a nil state?")
		}

		//set failure of child as root
		da.addFailure(temp, da.root)

		//create output
		da.addOutput(temp)
		q.Push(temp)
	}

	//step 2:do bfs to set failure
	for !q.IsEmpty() {
		n, ok := q.Pop()
		if !ok {
			return errors.New("An non-empty queue return a nil")
		}
		curr, ok := n.(*state)
		if !ok {
			return errors.New("It seems we insert a nil state?")
		}

		//fmt.Println("id", string(curr.c), curr.id)
		for i := curr.success.Begin(); i.LessEqual(curr.success.End()); i.Next() {
			entry := i.GetData()
			if entry == nil {
				return errors.New("An non-empty state has children but return nil")
			}
			transition := entry.GetKey().(rune)

			//calculate failure state
			temp, err := curr.nextState(transition)
			if err != nil {
				return err
			}

			//find nextState until nextState is not nil
			//if fail become root and the next state of root  is nil
			//return root
			fail := curr.failure
			for {
				next, err := fail.nextState(transition)
				if err != nil {
					return err
				} else if next == nil {
					fail = fail.failure
				} else {
					break
				}
			}
			fail, err = fail.nextState(transition)
			if err != nil {
				return nil
			}

			//add emit
			if err := temp.addEmits(fail.emits); err != nil {
				return err
			}
			//set failure
			da.addFailure(temp, fail)

			//set output
			if err := da.addOutput(temp); err != nil {
				return err
			}

			q.Push(temp)
		}
	}

	return nil
}

//addoupt put create a output array from state
func (da *Dart) addFailure(s *state, fail *state) {
	s.failure = fail
	da.fail[s.id] = fail.id
}

//addoupt put create a output array from state
func (da *Dart) addOutput(s *state) error {
	output := make([]int, 0)
	for i := s.emits.Begin(); i.LessEqual(s.emits.End()); i.Next() {
		v := i.GetData()
		if v == nil {
			return errors.New("An non-empty set return nil")
		}
		output = append(output, int(v.(value.Int)))
	}
	da.output[s.id] = output
	return nil
}

//Mactch return true if there exists pattern
func (da *Dart) Matches(text string) bool {
	b := []rune(text)
	curr := 0
	for i := range b {
		for !da.nextState(curr, b[i]) {
			curr = da.fail[curr]
		}
		curr = da.goTo(curr, b[i])
		if len(da.output[curr]) > 0 {
			return true
		}
	}
	return false
}

//nextState return id of next state accoding to currentstate and next char
func (da *Dart) nextState(currState int, char rune) bool {
	//find currState begin value and id
	begin := da.base[currState]
	id := begin + int(char)
	//check next state
	if id >= len(da.check) {
		return false
	} else if da.check[id] == begin {
		return true
	} else if da.check[id] != begin && currState == 0 {
		return true
	} else {
		return false
	}
}

//goTo return next state
func (da *Dart) goTo(currState int, char rune) int {
	//find currState begin value and id
	begin := da.base[currState]
	id := begin + int(char)
	//check next state
	if id > len(da.check) {
		return 0
	} else if da.check[id] == begin {
		return id
	} else {
		return 0
	}
}

//parsetext parse a text to find matched value
func (da *Dart) ParseText(text io.RuneReader, callback func([]int)) error {
	curr := 0
	for {
		//read a rune
		r, _, err := text.ReadRune()
		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF {
			return nil
		}

		for !da.nextState(curr, r) {
			curr = da.fail[curr]
		}

		curr = da.goTo(curr, r)

		if len(da.output[curr]) > 0 {
			if callback != nil {
				callback(da.output[curr])
			}
		}
	}

	return nil
}

//ParseString parse a string to find matched value
func (da *Dart) ParseString(text string, callback func([]int)) error {
	b := []rune(text)
	curr := 0
	for i := range b {
		for !da.nextState(curr, b[i]) {
			curr = da.fail[curr]
		}

		curr = da.goTo(curr, b[i])

		if len(da.output[curr]) > 0 {
			if callback != nil {
				callback(da.output[curr])
			}
		}
	}
	return nil
}
