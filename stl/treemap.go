package stl

//TreeMap is a RBTree
type TreeMap struct {
	r *RBTree
}

func NewTreeMap() *TreeMap {
	return &TreeMap{
		r: NewRBTree(),
	}
}

//Put will add a key into TreeMap if  not exist or update a key if exists
func (t *TreeMap) Put(e Entry) {
	//do only once
	t.r.lazyInit(e.(Value))

	//add a new node success
	if n, ok := t.r.add(e.(Value)); ok {
		t.r.size++
		t.r.balance(n)
	} else {
		n.val = e.(Value)
	}

}

//Delete will delete a key if exists or do nothing
func (t *TreeMap) Delete(e Entry) {
	t.r.Pop(e.(Value))
}

//Get will return value if exists a key or nil if  not exist
func (t *TreeMap) Get(e Entry) (interface{}, bool) {
	//find a key
	if n := t.r.find(e.(Value)); n != nil {
		return n.val.(Entry).GetValue(), true
	} else {
		return nil, false
	}
}

//Iterator will return a Iterator
func (t *TreeMap) Iterator(isAsc bool) func() (Entry, bool) {
	next := t.r.Iterator(isAsc)
	return func() (Entry, bool) {
		n, ok := next()
		return n.(Entry), ok
	}

}

//Len return the size of rbtree
func (t *TreeMap) Len() int {
	return t.r.Len()
}
