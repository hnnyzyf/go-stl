package stl

import (
	"errors"
	"sync"
)

type Value interface {
	Less(Value) bool
	More(Value) bool
}

const (
	RED = iota
	BLACK
)

type node struct {
	val Value

	color int

	l *node //l node
	r *node //r node
	p *node //parent node
}

func newnode(val Value) *node {
	return &node{
		val:   val,
		color: RED,
	}
}

func (n *node) clear() {
	n.val = nil
	n.color = RED
	n.l = nil
	n.r = nil
	n.p = nil
}

var rbPool = &sync.Pool{
	New: func() interface{} {
		return &node{
			color: RED,
		}
	},
}

func isBlack(n *node) bool {
	if n == nil {
		return true
	}

	if n.color == BLACK {
		return true
	}

	return false
}

func isRed(n *node) bool {
	if n == nil {
		return false
	}

	if n.color == RED {
		return true
	}

	return false
}

//brother will return a brother node
func (n *node) brother() *node {
	if n.p == nil {
		return n
	}

	if n == n.p.l {
		return n.p.r
	} else {
		return n.p.l
	}
}

//uncle will return a uncle node
func (n *node) uncle() *node {
	pa := n.p
	gp := pa.p
	if pa == gp.l {
		return gp.r
	} else {
		return gp.l
	}
}

//child will return a child node
func (n *node) child() *node {
	if n.l == nil {
		return n.r
	} else {
		return n.l
	}
}

type RBTree struct {
	Root *node
}

//a new RBTree
func NewRBTree() *RBTree {
	return &RBTree{}
}

//init
func (r *RBTree) init(val Value) {
	r.Root = &node{
		val:   val,
		color: BLACK,
	}
}

//push a new node
func (r *RBTree) Push(val Value) {
	//do only once
	if r.Root == nil {
		r.init(val)
	}

	//add a new node success
	if n, ok := r.add(val); ok {
		r.Balance(n)
	}

}

//pop will remove a node
func (r *RBTree) Pop(val Value) {
	if r.Root == nil {
		return
	}

	//find the node which will be deleted
	if n := r.find(val); n != nil {
		//find the node which will be deleted
		del := r.replace(n)
		//exchange the Value
		n.val = del.val
		//delete the node and get the child node
		c, ok := r.delete(del)
		if !ok {
			r.ReBalance(c)
		}
	}

}

//Balance will make a RBTree become balance
func (r *RBTree) Balance(n *node) {
	//begin
	curr := n
	for {
		if curr == r.Root {
			//case 1:新节点N位于树的根上，没有父节点。在这种情形下，我们把它重绘为黑色以满足性质2
			curr.color = BLACK
			break

		} else if curr.p.color == BLACK {
			//Case 2:新节点的父节点P是黑色
			break

		} else if uc := curr.uncle(); uc != nil && uc.color == RED {
			//Case 3:p and uncle is red
			pa := curr.p
			uc := curr.uncle()
			gp := pa.p

			//change color
			pa.color = BLACK
			uc.color = BLACK
			gp.color = RED
			//continue Balance，set grandp node as current node to rebalance
			curr = gp
			continue

		} else if bro := curr.brother(); bro != nil && bro.color == BLACK || bro == nil {
			//case 4,5:p is red and uncle is nil or black
			pa := curr.p
			gp := pa.p

			if pa.r == curr && gp.l == pa {
				//case 4:curr is r node pf p and p is the l node of g
				r.r2l(curr)
				curr = pa
			} else if pa.l == curr && gp.r == pa {
				//case 4:curr is l node pf p and p is the r node of gp
				r.r2r(curr)
				curr = pa
			} else if curr == pa.l && pa == gp.l {
				//case 5:curr is r node pf p and p is the r node of gp
				pa.color = BLACK
				gp.color = RED
				r.r2r(pa)
				break
			} else if curr == pa.r && pa == gp.r {
				//case 5:curr is r node pf p and p is the r node of gp
				pa.color = BLACK
				gp.color = RED
				r.r2l(pa)
				break
			}

		} else {
			panic("Impossible state")
		}
	}
}

//rebalance will be executed alter a black node is deleted
func (r *RBTree) ReBalance(n *node) {
	curr := n
	for {

		if isBlack(curr) {
			if curr == nil || curr == r.Root {
				//case 0:n is nil
				break
			} else if pa := curr.p; pa.r == curr && isRed(curr.r) {
				//case 1:one child node is red and have the same direction
				curr.r.color = BLACK
				curr.color = pa.color
				pa.color = BLACK

				r.r2l(curr)
				break
			} else if pa := curr.p; pa.l == curr && isRed(curr.l) {
				//case 1:one child node is red and have the same direction
				curr.l.color = BLACK
				curr.color = pa.color
				pa.color = BLACK

				r.r2r(curr)
				break
			} else if pa := curr.p; isBlack(curr.l) && isBlack(curr.r) {
				//case 2: both child nodes are black
				curr.color = RED
				if pa.color == RED {
					pa.color = BLACK
					break
				} else {
					pa.color = BLACK
					curr = pa.brother()
				}
			} else if pa := curr.p; pa.r == curr && isRed(curr.l) {
				//case 3: one child is red and have the oposite direction
				c := curr.l
				curr.color = RED
				c.color = BLACK
				//rotate
				r.r2r(c)
				//rebalance
				curr = c
			} else if pa := curr.p; pa.l == curr && isRed(curr.r) {
				//case 3: one child is red and have the oposite direction
				c := curr.r
				curr.color = RED
				c.color = BLACK
				//rotate
				r.r2l(c)
				//rebalance
				curr = c
			} else {
				panic("Impossible state")
			}
		} else if pa := curr.p; isRed(curr) {
			curr.color = BLACK
			pa.color = RED
			if pa.l == curr {
				r.r2r(curr)
				curr = pa.l
			} else {
				r.r2l(curr)
				curr = pa.r
			}
		} else {
			panic("Impossible state")
		}
	}
}

func (r *RBTree) r2l(n *node) {
	pa := n.p
	gp := pa.p

	pa.r = n.l
	if n.l != nil {
		n.l.p = pa
	}

	n.l = pa
	pa.p = n

	if gp == nil {
		r.Root = n
	} else if gp.l == pa {
		gp.l = n
	} else {
		gp.r = n
	}
	n.p = gp
}

func (r *RBTree) r2r(n *node) {
	pa := n.p
	gp := pa.p

	pa.l = n.r
	if n.r != nil {
		n.r.p = pa
	}

	n.r = pa
	pa.p = n
	if gp == nil {
		r.Root = n
	} else if gp.l == pa {
		gp.l = n
	} else {
		gp.r = n
	}
	n.p = gp

}

//add will add a new node
func (r *RBTree) add(val Value) (*node, bool) {
	curr := r.Root
	for {

		switch {
		//smaller than root then turn to l
		case val.Less(curr.val):
			if curr.l != nil {
				curr = curr.l
			} else {
				//add new node
				curr.l = rbPool.Get().(*node)
				curr.l.val = val
				//curr.l =newnode(val)
				curr.l.p = curr
				return curr.l, true
			}
		//bigger than root then turn to r
		case val.More(curr.val):
			if curr.r != nil {
				curr = curr.r
			} else {
				//add new node
				curr.r = rbPool.Get().(*node)
				curr.r.val = val
				//curr.r=newnode(val)
				curr.r.p = curr
				return curr.r, true
			}
		default:
			return nil, false
		}

	}
}

//delete will find the node
func (r *RBTree) find(val Value) *node {
	curr := r.Root
	for {
		switch {
		//smaller than root then turn to l
		case val.Less(curr.val):
			if curr.l != nil {
				curr = curr.l
			} else {
				return nil
			}
		//bigger than root then turn to r
		case val.More(curr.val):
			if curr.r != nil {
				curr = curr.r
			} else {
				return nil
			}
		default:
			return curr
		}
	}
}

//replace will find the node
func (r *RBTree) replace(n *node) *node {

	if n.r != nil {
		curr := n.r
		for curr.l != nil {
			curr = curr.l
		}
		return curr
	}

	if n.l != nil {
		curr := n.l
		for curr.r != nil {
			curr = curr.r
		}
		return curr
	}

	return n
}

//delete will delete a node
func (r *RBTree) delete(n *node) (*node, bool) {
	if n == r.Root {
		r.Root = nil
		return nil, true
	}

	pa := n.p
	bro := n.brother()
	//case 1:l node
	if n.l == nil && n.r == nil {
		if pa.l == n {
			pa.l = nil
		} else {
			pa.r = nil
		}
		if n.color == RED {
			n.clear()
			rbPool.Put(n)
			return nil, true
			//if a leaf node is black,it must have brother node
		} else {
			return bro, false
		}
	}

	//have a l node
	if n.l != nil && n.r == nil {
		if pa.l == n {
			pa.l = n.l
		} else {
			pa.r = n.l
		}
		n.l.p = pa

		if n.l.color == RED {
			n.l.color = BLACK
			n.clear()
			rbPool.Put(n)
			return nil, true
			//if child is black,we need to do a reblance
		} else {
			return bro, false
		}
	}

	//have a r node
	if n.r != nil && n.l == nil {
		if pa.r == n {
			pa.r = n.r
		} else {
			pa.l = n.r
		}
		n.r.p = pa

		if n.r.color == RED {
			n.r.color = BLACK
			n.clear()
			rbPool.Put(n)
			return nil, true
			//if child is black,we need to do a reblance
		} else {
			return bro, false
		}
	}

	return nil, true

}

type Hook func(n *node) bool

//do bfs for a RBtree
func (r *RBTree) BFS(hook Hook) bool {
	if r.Root == nil {
		return true
	}

	queue := []*node{r.Root}
	for idx := 0; idx < len(queue); idx++ {
		n := queue[idx]
		if hook != nil && !hook(n) {
			return false
		}

		if n.l != nil {
			queue = append(queue, n.l)
		}

		if n.r != nil {
			queue = append(queue, n.r)
		}

	}

	return true
}

//test a RBTree
func (r *RBTree) IsRBTree(testRoot Hook, testRedNode Hook, testPath Hook) (bool, error) {

	if r.Root == nil {
		return true, nil
	}

	//性质2
	if !testRoot(r.Root) {
		return false, errors.New("The root is RED")
	}

	//性质4
	if !r.BFS(testRedNode) {
		return false, errors.New("The RED node does not have 2 BLACK nodes")
	}

	//性质5
	if !r.BFS(testPath) {
		return false, errors.New("The Path is not equal")
	}

	return true, nil
}
