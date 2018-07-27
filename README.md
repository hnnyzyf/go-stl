# go-stl 

This package contains differnt data structures which golang does not provide.

Package container provides following structures :
  
 - [Deque](/doc/deque.md)
 - [Stack](/doc/stack.md)
 - [Queue](/doc/queue.md)
 - [RBtree](/doc/rbtree.md) 
 - [TreeMap](/doc/treemap.md)
 - [TreeSet](/doc/treeset.md)
 - [Value](/doc/value.md)
 - [Pair](/doc/pair.md)



Package algorithm provides following algorithms :

  - [Kmp](/doc/kmp.md)
  - [Manacher](/doc/manacher.md)
  - [Dart (Static AhoCorasickDoubleArrayTrie)](/doc/dart.md)
  - [Cedra (Dynamic AhoCorasickDoubleArrayTrie)](/doc/cedra.md)



## Import and create a instance

Package container provides following structures :
  - deque
  ```
  import "github.com/hnnyzyf/go-stl/container/deque"

  d:=deque.New()
  ```
  - stack
  ```
  import "github.com/hnnyzyf/go-stl/container/stack"

  s:=stack.New()
  ```
  - queue
   ```
  import "github.com/hnnyzyf/go-stl/container/queue"

  q:=queue.New()
  ```
  - rbtree
  ```
  import "github.com/hnnyzyf/go-stl/container/rbtree"

  r:=rbtree.New()
  ```
  - treemap
  ```
  import "github.com/hnnyzyf/go-stl/container/treemap"

  tm:=treemap.New()
  ```
  - treeset
  ```
  import "github.com/hnnyzyf/go-stl/container/treeset"

  ts:=treeset.New()
  ```
  - pair
  ```
  import "github.com/hnnyzyf/go-stl/container/pair"

  s:=pair.String()
  i:=pair.Int()
  ui:=pair.Uint64()
  f:=pair.Floate64()
  r:=pair.Rune()

  ```
  - value
   ```
  import "github.com/hnnyzyf/go-stl/container/value"

  s:=value.String()
  i:=value.Int()
  ui:=value.Uint64()
  f:=value.Floate64()
  r:=value.Rune()
  ```

Package algorithm provides following algorithms :

  - kmp
  ```
  import "github.com/hnnyzyf/go-stl/algorithm/kmp"

  text:="abcdef"
  k:=kmp.New(text)

  ```
  - manacher
  ```
  import "github.com/hnnyzyf/go-stl/algorithm/manacher"

  start:=`$`
  delimiter:=`#`
  m:=manacher.New(start,delimiter)

  ```
  - dart (Static AhoCorasickDoubleArrayTrie)
  ```
  import "github.com/hnnyzyf/go-stl/algorithm/dart"

  d:=dart.New()

  ```
  - cedra (Dynamic AhoCorasickDoubleArrayTrie)
  ```
  import "github.com/hnnyzyf/go-stl/algorithm/cedra"

  c:=cedra.New()

  ```

## API
Package container provides following variable :

  - value
```
type Value interface {
	Less(Value) bool
	More(Value) bool
	Equal(Value) bool
}
```
  - pair

```
type Pair interface {
	value.Value
	GetKey() interface{}
	GetValue() interface{}
}

```

Package container provides following structures :

  - deque
```
func (d *Deque) PushFront(val interface{})
func (d *Deque) PushBack(val interface{})
func (d *Deque) PopFront() (interface{}, bool)
func (d *Deque) PopBack() (interface{}, bool)
func (d *Deque) Len() int
func (d *Deque) IsEmpty() bool

```
  - stack
```
func (s *Stack) Push(val interface{})
func (s *Stack) Pop() (interface{}, bool)
func (s *Stack) Top() interface{}
func (s *Stack) IsEmpty() bool 
func (s *Stack) Len() int
```
  - queue
```
func (q *Queue) Push(val interface{})
func (q *Queue) Pop() (interface{}, bool) 
func (q *Queue) Front() interface{}
func (q *Queue) Back() interface{}
func (q *Queue) IsEmpty() bool
func (q *Queue) Len() int 
```
  - rbtree
```
func (r *RBTree) Push(val value.Value)
func (r *RBTree) Pop(val value.Value)
func (r *RBTree) Find(val value.Value) value.Value 
func (r *RBTree) Get(i *Riterator) value.Value
func (r *RBTree) Len() int
```
  - treemap
```
func (t *TreeMap) Insert(e pair.Pair)
func (t *TreeMap) Erase(e pair.Pair)
func (t *TreeMap) Find(e pair.Pair) pair.Pair 
func (t *TreeMap) Get(i *Miterator) pair.Pair 
func (t *TreeMap) Len() int 
func (t *TreeMap) IsEmpty() bool
```
  - treeset
```
func (s *TreeSet) Insert(val value.Value) 
func (s *TreeSet) Erase(val value.Value) 
func (s *TreeSet) Find(val value.Value) bool
func (s *TreeSet) Len() int 
```

## Iterator

Package container provides following structures :
  - deque
```
  d:=deque.New()
  for i:=0;i<d.Len();i++{
      v:=d.Get(i)
  }

```
  - rbtree
```
  r:=rbtree.New()
  //Forward
  for i:=r.Begin();i.LessEqual(r.End());i.Next(){
      v:=i.GetData()
  }

  //Reverse
  for i:=r.End();i.MoreEqual(r.Begin());i.Last(){
      v:=i.GetData()
  }
```
  - treemap
```
  tm:=treemap.New()
  //Forward
  for i:=tm.Begin();i.LessEqual(tm.End());i.Next(){
       v:=i.GetData() 
    or v:=tm.Get(i)
  }

  //Reverse
  for i:=tm.End();i.MoreEqual(tm.Begin());i.Last(){
       v:=i.GetData()
    or v:=tm.Get(i)
  }

```
  - treeset
```
  ts:=treeset.New()
  //Forward
  for i:=ts.Begin();i.LessEqual(ts.End());i.Next(){
       v:=i.GetData()
    or v:=ts.Get(i)
  }

  //Reverse
  for i:=ts.End();i.MoreEqual(ts.Begin());i.Last(){
       v:=i.GetData()
    or v:=ts.Get(i)
  }

```

## Benchmark

 - deque
    
```
goos: linux
goarch: amd64
BenchmarkPushfront  20000000           158 ns/op
BenchmarkPushback   20000000           163 ns/op
BenchmarkPush       10000000           218 ns/op

```
 - stack

 ```
goos: linux
goarch: amd64
BenchmarkPushMix    20000000           176 ns/op
BenchmarkMix        10000000           196 ns/op

```
 - rbtree

 ```
goos: linux
goarch: amd64
BenchmarkPush        2000000           699 ns/op
BenchmarkMix         1000000          1054 ns/op
BenchmarkGet         1000000          1068 ns/op
BenchmarkIterator   10000000          1087 ns/op

```
 - treemap

  ```
goos: linux
goarch: amd64
BenchmarkPut     2000000           919 ns/op
BenchmarkGet     1000000          1111 ns/op

```

 - treeset

  ```
goos: linux
goarch: amd64
BenchmarkPush        2000000           699 ns/op
BenchmarkMix         1000000          1054 ns/op
BenchmarkGet         1000000          1068 ns/op
BenchmarkIterator   10000000          1087 ns/op

```

