## 前言

stack 是一个非常常用的数据结构，我们使用deque来作为底层数据结构封装stack


### Stack的操作和基本数据结构

栈的操作只有两种

- Push(入栈)
- Pop(出栈)


基本数据结构，很明显栈就是一个deque

```
type Stack struct {
    d *Deque
}
```


#### 入栈

封装了deque的pushback操作


```
func (s *Stack) Push(val interface{}) {
    s.d.PushBack(val)
}
```



#### 出栈

封装了 deque 的popback操作


```
func (s *Stack) Pop() (interface{}, bool) {
    return s.d.PopBack()
}
```
