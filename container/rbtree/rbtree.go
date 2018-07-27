package rbtree

import (
	"errors"
	"github.com/hnnyzyf/go-stl/container/value"
)

const (
	RED = iota
	BLACK
)

type node struct {
	val value.Value

	color int

	l *node //l node
	r *node //r node
	p *node //parent node
}

func newnode(val value.Value) *node {
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
	root *node
	size int
}

//a new RBTree
func New() *RBTree {
	return &RBTree{}
}

//init
func (r *RBTree) lazyInit(val value.Value) {
	if r.root == nil {
		r.root = &node{
			val:   val,
			color: BLACK,
		}
		r.size++
	}
}

//push a new node
func (r *RBTree) Push(val value.Value) {
	//do only once
	r.lazyInit(val)

	//add a new node success
	if n, ok := r.add(val); ok {
		r.size++
		r.balance(n)
	}
}

//pop will remove a node
func (r *RBTree) Pop(val value.Value) {
	if r.root == nil {
		return
	}

	//find the node which will be deleted
	if n := r.find(val); n != nil {
		//find the node which will be deleted
		del := r.replace(n)
		//exchange the value.Value
		n.val = del.val
		//delete the node and get the child node
		c, ok := r.delete(del)
		r.size--
		if !ok {
			r.rebalance(c)
		}
	}

}

//Find will return a value.Value
func (r *RBTree) Find(val value.Value) value.Value {
	//when no node to find
	if r.root == nil{
		return nil
	}

	//have at leaset one node
	if n := r.find(val); n != nil {
		return n.val
	} else {
		return nil
	}
}

//Get will return a value.Value
func (r *RBTree) Get(i *Riterator) value.Value {
	if i.n != nil {
		return i.n.val
	} else {
		return nil
	}
}

func (r *RBTree) Begin() *Riterator {
	if r.root == nil {
		return newRiterator(nil, r.root)
	}

	curr := r.root
	for curr.l != nil {
		curr = curr.l
	}
	return newRiterator(curr, r.root)
}

func (r *RBTree) End() *Riterator {
	if r.root == nil {
		return newRiterator(nil, r.root)
	}

	curr := r.root
	for curr.r != nil {
		curr = curr.r
	}
	return newRiterator(curr, r.root)
}

//Len will return size of rbtree
func (r *RBTree) Len() int {
	return r.size
}

//balance will make a RBTree become balance
func (r *RBTree) balance(n *node) {
	//begin
	curr := n
	for {
		if curr == r.root {
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
			//continue balance，set grandp node as current node to rebalance
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
func (r *RBTree) rebalance(n *node) {
	curr := n
	for {

		if isBlack(curr) {
			if curr == nil || curr == r.root {
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
		r.root = n
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
		r.root = n
	} else if gp.l == pa {
		gp.l = n
	} else {
		gp.r = n
	}
	n.p = gp

}

//add will add a new node
func (r *RBTree) add(val value.Value) (*node, bool) {
	curr := r.root
	for {

		switch {
		//smaller than root then turn to l
		case val.Less(curr.val):
			if curr.l != nil {
				curr = curr.l
			} else {
				//add new node
				//curr.l = rbPool.Get().(*node)
				//curr.l.val = val
				curr.l = newnode(val)
				curr.l.p = curr
				return curr.l, true
			}
		//bigger than root then turn to r
		case val.More(curr.val):
			if curr.r != nil {
				curr = curr.r
			} else {
				//add new node
				//curr.r = rbPool.Get().(*node)
				//curr.r.val = val
				curr.r = newnode(val)
				curr.r.p = curr
				return curr.r, true
			}
		default:
			curr.val = val
			//if equal,return the node
			return curr, false
		}

	}
}

//delete will find the node
func (r *RBTree) find(val value.Value) *node {
	curr := r.root
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
	if n == r.root {
		r.root = nil
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
			//rbPool.Put(n)
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
			//rbPool.Put(n)
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
			//rbPool.Put(n)
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
	if r.root == nil {
		return true
	}

	queue := []*node{r.root}
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
func (r *RBTree) IsRBTree(testroot Hook, testRedNode Hook, testPath Hook) (bool, error) {

	if r.root == nil {
		return true, nil
	}

	//性质2
	if !testroot(r.root) {
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

//the iterator of RBTree
type Riterator struct {
	n    *node
	root *node
	id   int
}

func newRiterator(n *node, root *node) *Riterator {
	return &Riterator{
		n:    n,
		root: root,
	}
}

//value.Value return the val in node
func (i *Riterator) GetData() value.Value {
	if i.n == nil {
		return nil
	} else {
		return i.n.val
	}
}

//mid order
func (i *Riterator) Next() {
	curr := i.n
	//have right node
	if curr == nil {
		//do nothing
	} else if curr.r != nil {
		curr = curr.r
		for curr.l != nil {
			curr = curr.l
		}
		//node does not have right node,next node is its parent
	} else if curr.p != nil && curr.p.l == curr {
		curr = curr.p
		//node does not have right node,next node is its parent
	} else if curr.p != nil && curr.p.r == curr {
		curr = curr.p
		for curr.p != nil && curr.p.r == curr {
			curr = curr.p
		}

		//left child of root
		if curr.p != nil {
			curr = curr.p
		} else {
			//right child of root
			curr = i.n
		}
	} else if curr.p == nil {
		//root does not have right node and right node is nil
		//donothing
	} else {
		panic("Impossible state")
	}
	if i.n == curr {
		i.n = nil
	} else {
		i.n = curr
	}
}

//back order
func (i *Riterator) Last() {
	curr := i.n
	//have left node//have right node
	if curr == nil {
		//do nothing
	} else if curr.l != nil {
		curr = curr.l
		for curr.r != nil {
			curr = curr.r
		}
		//node does not have right node,next node is its parent
	} else if curr.p != nil && curr.p.r == curr {
		curr = curr.p
		//node does not have right node,next node is its parent
	} else if curr.p != nil && curr.p.l == curr {
		curr = curr.p
		for curr.p != nil && curr.p.l == curr {
			curr = curr.p
		}

		//left child of root
		if curr.p != nil {
			curr = curr.p
		} else {
			//right child of root
			curr = i.n
		}
	} else if curr.p == nil {
		//root does not have right node and right node is nil
		//donothing
	} else {
		panic("Impossible state")
	}

	if i.n == curr {
		i.n = nil
	} else {
		i.n = curr
	}
}

func (i *Riterator) Equal(o *Riterator) bool {
	return i.n == o.n
}

func (i *Riterator) LessEqual(o *Riterator) bool {
	if i.n != nil && o.n != nil && (!i.n.val.More(o.n.val)) {
		return true
	} else {
		return false
	}
}

func (i *Riterator) MoreEqual(o *Riterator) bool {
	if i.n != nil && o.n != nil && (!i.n.val.Less(o.n.val)) {
		return true
	} else {
		return false
	}
}
