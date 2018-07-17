package stl

import (
	"errors"
)

type node struct {
	base  int
	check int
	value []interface{}
}

type block struct {
	prev   int
	next   int
	num    uint8
	reject uint8
	trial  int
	ehead  int
}

type info struct {
	sibling rune
	child   rune
}

type Cedra struct {
	nodes []*node

	blocks []*block

	infos []*info

	IsOrder  bool
	MaxTrial int
}

func NewCedra(IsOrder bool, MaxTrial int) *Cedra {
	return &Cedra{
		IsOrder:  IsOrder,
		MaxTrial: MaxTrial,
	}
}

//Build creates a dynamic double array tries
func (da *Cedra) Build(keys []string) error {
	for i := range keys {
		return da.insert(keys[i])
	}
}

//insert will add a key
func (da *Cedra) insert(key string) error {
	return da.update(key, 0, 0)
}

//update will update a key from pos to end and add it alter the from state
func (da *Cedra) update(key string, from int, pos int) error {
	b := []rune(key)[pos:]

	if len(b) <= 0 {
		return errors.New("failed to insert zero-length key")
	}

	offset := from >> 32
	if offset != 0 {
	}
}

//follow return the next state after curr state
func (da *Cedra) follow(from int, lable rune) int {
	base := da.nodes[from].base
	to := 0
	if to = base ^ label; base < 0 || da.nodes[to].check < 0 {
		to = da.popEnode(base, lable, from)
		da.pushSibling(from, to^lable, lable, base >= 0)
	} else if da.nodes[to].check != from {
		to = da.resolve(from, base, lable)
	} else {
		//do nothing
	}
	return to
}

//popenode pop empty node from block; never transfer the special block (bi = 0)
func (da *Cedra) popEnode(base int, lable rune, from int) int {
	e := base ^ lable
	if e >= 0 {
		e = da.findPlace()
	}
	bi := e >> 8

	n := da.nodes[e]
	b := da.blocks[bi]
	b.num--
	if b.num == 0 {
		if bi != 0 {
			da.transferBlock(bi)
		}
	} else {
		da.nodes[-n.base].check = n.check
		da.nodes[-n.check].base = n.base

		if e == b.ehead {
			b.ehead = -n.check
		}

		if bi == 0 && b.num == 1 && b.trial != da.MaxTrial {
			da.transferBlock(bi)
		}
	}

	// initialize the released node
	if lable != 0 {
		n.base = -1
	} else {
		n.value = 0
	}

	n.check = from

	if base < 0 {
		da.nodes[from].base = e ^ lable
	}
	return e
}

func (da *Cedra) pushSibling(from int, base int, lable rune, flag bool) {
	r := &da.infos[from].child
	if flag && ((da.IsOrder && lable > *r) || (!da.IsOrder && *r == 0)) {
		for da.IsOrder && *r != 0 && *r < label {
			r = &da.infos[base^*r].sibling
		}
	}

	*r = lable
	da.infos[base^lable].sibling = *r
}

// resolve conflict on base_n ^ label_n = base_p ^ label_p
func (da *Cedra) resolve(nfrom int, nbase int, nlable rune) int {
	// examine siblings of conflicted nodes
	pto := nbase ^ nlable
	pfrom := da.nodes[pto].check
	pbase := da.nodes[pfrom].base

	// whether to replace siblings of newly added
	flag := da.consult(nbase, pbase, da.infos[nfrom].child, da.infos[pfrom].child)

	var children []rune
	if flag {
		children = da.setChild(nbase, da.infos[nfrom].child, int(nlabel))
	} else {
		children = da.setChild(pbase, da.infos[pfrom].child, -1)
	}

	first := children[0]
	last := children[len(children)-1]

	var base int
	if first == last {
		base = da.findPlace()
	} else {
		base = da.findPlace(first, last) ^ first
	}

	// replace & modify empty list
	var from int
	var base_ int
	if flag {
		from = nfrom
		base_ = nbase
	} else {
		from = pfrom
		base_ = pbase
	}

	// new child
	if flag && first == nlable {
		da.infos[from].child = nlable
	}
	// new base
	da.nodes[from].base = base
	// to_ => to
	for i := range children {
		to := da.popEnode(base, children[i], from)
		to_ := base_ ^ children[i]
		if i == len(children)-1 {
			da.infos[to].sibling = 0
		} else {
			da.infos[to].sibling = children[i+1]
		}

		// skip newcomer (no child)
		if flag && to_ == pto {
			continue
		}

		//cf?

		n := da.nodes[to]
		n_ := da.nodes[to_]

		// copy base; bug fix
		if n.base = n_.base; n.base > 0 && children[i] != 0 {
			da.infos[to].child = da.infos[to_].child
			c := _ninfo[to].child

		}

	}

}

// enumerate (equal to or more than one) child nodes
func (da *Cedra) setChild(base int, child rune, lable int) []rune {
	children := make([]rune)

	if child == 0 {
		children = append(children, child)
		child = da.infos[base^child].sibling
	}

	if da.IsOrder {
		for child != 0 && child < lable {
			children = append(children, child)
			child = da.infos[base^child].sibling
		}
	}

	if lable != -1 {
		children = append(children, rune(lable))
	}

	for child != 0 {
		children = append(children, child)
		child = da.infos[base^child].sibling
	}

	return children
}

func (da *Cedra) Open() {}

func (da *Cedra) Save() {}
