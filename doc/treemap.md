# 前言

treemap 是一个由红黑树实现的map,可以提供平均O(logn)的时间复杂度的查找

同时可以提供顺序的key的遍历，包括正向的(使用中序遍历)和反向的(使用后续遍历)

### treemap的数据结构

##### treemap是一颗红黑树

```
//TreeMap is a RBTree
type TreeMap struct {
    r *RBTree
}
```

##### 数据域是key和value构成的entry

Entry(java中习惯用Entry)有一个别名是Pair(c++中习惯用pair)
```
type Pair = Entry

type Entry interface {
    Value
    GetKey() interface{}
    GetValue() interface{}
}
```

### treemap的基本操作

- Put 增添一个Entry 或者更新一个key

```
func (t *TreeMap) Put(e Entry)
```

- Delete 删除一个Entry

```
func (t *TreeMap) Delete(e Entry)
```

- Get 获得一个Entry的value

```
//具体的类型需要手动转换
func (t *TreeMap) Get(e Entry) (interface{}, bool)
```

- Iterator 返回一个迭代器


```
//迭代器是一个函数闭包
func (t *TreeMap) Iterator(isAsc bool) func() (Entry, bool)
```

#### treemap迭代器

treemap的迭代器也是会失效的，如果再迭代的过程中插入或者删除某个Entry,当然，修改Entry是不会让迭代器失效的


##### 正向遍历

正向遍历的过程是中序遍历的过程，我们将每次中旬遍历封装成函数闭包，这样我们就可以实现迭代的遍历了

正向遍历 封装了rbtree的正向遍历

```
//Next will return a function closure
func (r *RBTree) asc() func() (Value, bool) {
    //create a Stack
    s := NewStack()
    
    //the first node is r.root
    curr := r.root
    
    //create a function closure
    next := func() (Value, bool) {
        //put the left child in the stack
        for curr != nil {
            s.Push(curr)
            curr = curr.l
        }
        
        //if curr is nil,pop stack
        if e, ok := s.Pop(); ok {
            n := e.(*node)
            curr = n.r
            return n.val, true
        } else {
            //only happen when all node have been visited
            return nil, false
        }
    }
    
    return next
}
```


##### 反向遍历

反向遍历的过程是后序遍历的过程,思路和正向遍历相同


反向遍历 封装了rbtree的反向遍历

```
//Next will return a function closure
func (r *RBTree) desc() func() (Value, bool) {
    //create a Stack and add the first element
    s := NewStack()
    curr := r.root
    next := func() (Value, bool) {
        for curr != nil {
            s.Push(curr)
            curr = curr.r
        }
        //if curr does not have a right child,pop stack
        if e, ok := s.Pop(); ok {
            n := e.(*node)
            curr = n.l
            return n.val, true
        } else {
            //only happen when all node have been visited
            return nil, false
        }
    }
    return next
}
```

