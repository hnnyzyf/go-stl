# go-stl 

This package contains differnt data structures which golang does not provide:
 
 - [Deque](/doc/deque.md)
    
```
goos: linux
goarch: amd64
BenchmarkPushfront  20000000           158 ns/op
BenchmarkPushback   20000000           163 ns/op
BenchmarkPush       10000000           218 ns/op

```
 - [Stack](/doc/stack.md)

 ```
goos: linux
goarch: amd64
BenchmarkPushMix    20000000           176 ns/op
BenchmarkMix        10000000           196 ns/op

```
 - [RBtree](/doc/rbtree.md) 

 ```
goos: linux
goarch: amd64
BenchmarkPush        2000000           699 ns/op
BenchmarkMix         1000000          1054 ns/op
BenchmarkGet         1000000          1068 ns/op
BenchmarkIterator   10000000          1087 ns/op

```
 - [TreeMap](/doc/treemap.md)

  ```
goos: linux
goarch: amd64
BenchmarkPut     2000000           919 ns/op
BenchmarkGet     1000000          1111 ns/op

```
 - [Entry](/doc/entry.md)

  ```
No Test

```
 - [acDAT](/doc/acdat.md)
