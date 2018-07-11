## 前言

在使用golang的过程中,stack是非常常用的数据结构，但是golang并没有提供原生的stack使用,因此参考stl中stack的实现方式，设计了stack的底层数据结构deque

## 数据结构和语言

STL中的deque设计了一块map的映射区域,映射区域每段存储一个数据缓冲区chunck

Deque使用golang实现

```
type Deque struct {
    mmap []chunck //map

    begin *iterator
    end   *iterator
}
```

数据缓冲区的定义

```
type chunck []interface{}
```

魔数的定义


```
const (
    MapSize    = 8
    ChunckSize = 128
)
```



### 迭代器的介绍

作为一个标准的容器，迭代器是必不可少的，golang中因为没有运算符重载,迭代器的操作只能用函数表示，看起来是丑了点，不容C++看起来那么简练

迭代器的定义

```
//iterator is the index of bucket
type iterator struct {
    chunck int
    index  int
}
```

我们在deque创建的时候，分别定义了两个迭代器指向deque数据的起始和终止位置


```
begin *iterator
end   *iterator
```

begin,end初始化为中间位置
```
begin: &iterator{
    chunck: (MapSize + 1) / 2,
    index:  0,
}
end: &iterator{
    chunck: (MapSize - 1) / 2,
    index:  ChunckSize - 1,
},
```





### deque的空间分配

#### map区域的大小
    
map区域初始的大小为 8 ，由常数Mapsize确定

如果map满了，我们使用如下策略进行扩容

- 小于1024的话,map每次扩容一倍
- 大于等于1024的话,map每次扩容25%

这个策略和golang 原生数组扩容的策略是一致的

```
//if size of d.mmap is smaller than 1024,we double every time
//if size of d.mmap is bigger than 1024,we add 25% every time
if len(d.mmap) < 1024 {
    mmap = make([]chunck, 2*len(d.mmap))
} else {
    mmap = make([]chunck, len(d.mmap)/4+len(d.mmap))
}
```

#### 数据缓冲区的大小

数据缓冲区的使用一个固定长度为128的接口切片数组实现,由常数chunckSize确定

因为接口内存储的实际是指针(64位服务器上8个字节)或者基本数据类型(整形，浮点数类型等等)

我们可以认为一个数据缓冲区的大小大最多为1024B

#### 数据缓冲区的创建时间

如果每次mmap扩容都创建新的数据缓冲取得话，可能会有很多空间的浪费，所以只有再新的数据缓冲区被使用时，我们才换创建一个新的数据缓冲区


### map的重分配策略

#### 存储空间不足时的重分配

分为如下三种情况:

 - Case1:尾部空间不足,头部空间剩余
 - Case2:尾部空间剩余,头部空间不足
 - Case3:头部尾部空间均不足

##### Case1：尾部空间不足,头部空间剩余

我们只需要调整mmap的begin和end的位置，将头部的部分空间匀给尾部即可,而不需要重新分配一块mmap的存储空间

因为尾部插满了，所以在移动的时候我们尽可能给尾部多分配点存储空间


```
if d.end.chunck == len(d.mmap)-1 && d.begin.chunck >= 1 {
    //cal the offset
    offset := (d.begin.chunck + 1) / 2
    //copy all between d.begin.chunck and d.end.chunck
    copy(d.mmap[d.begin.chunck-offset:d.end.chunck+1-offset], d.mmap[d.begin.chunck:d.end.chunck+1])

    //set nil
    for i := range d.mmap[len(d.mmap)-offset:] {
        d.mmap[i+len(d.mmap)-offset] = nil
    }

    //reindex
    d.begin.chunck -= offset
    d.end.chunck -= offset

    //begin has no space and end has space
}
```

##### Case2：尾部空间剩余,头部空间不足

这种情况和case1是一样的，只不过是把尾部的存储空间分配给头部



```
if d.begin.chunck == 0 && d.end.chunck <= len(d.mmap)-2 {

    //cal the offset
    offset := (len(d.mmap) - d.end.chunck) / 2

    //copy all between d.begin.chunck and d.end.chunck
    copy(d.mmap[offset:d.end.chunck+offset+1], d.mmap[:d.end.chunck+1])

    //set nil
    for i := range d.mmap[:offset] {
        d.mmap[i] = nil
    }

    //reindex
    d.begin.chunck += offset
    d.end.chunck += offset

}
```

##### case3：头部尾部空间均不足

这个时候，我们必须申请一块新的存储空间给mmap了，并将老的mmap的数据拷贝到新的mmap中,因为mmap本身并没有多大，这个拷贝相对来书还是比较廉价的



#### 存储空间剩余太多的时候的重分配

其实后多数据结构是不释放已经分配的空间的，比如golang原生的map,delete的空间只是标记掉但是没有释放掉，所以一直在思考是否需要释放大量不用的空间，考虑再三，决定还是要做这一块，毕竟内存是很宝贵的,能省即省

回收内存有两种情况:

 - Case 1:chunck中的元素已经全部被删除掉了
 - Case 2:mmap中，未使用的空间比已经使用的空间还要大


##### Case 1:chunck中的元素已经全部被删除掉了
在执行 Pop操作的时候，有时一个chunck钟的全部数据已经被删除掉了，我们要执行一次chunk的回收操作,为了尽可能减少对象的产生

我们使用一个全局的sync.Pool 池来回收这些chunck,因为很有可能马上又要创建一个chunck，这样我们就能节省一次内存分配


```
//if old chunck does not been used ,revoke it
    if chunck != d.end.chunck {
        dequePool.Put(d.mmap[d.end.chunck])
        d.mmap[d.end.chunck] = nil
    }


//if old chunck does not been used ,revoke it
    if chunck != d.begin.chunck {
        dequePool.Put(d.mmap[d.begin.chunck])
        d.mmap[d.begin.chunck] = nil
    }
```

##### Case 2:mmap中，未使用的空间比已经使用的空间还要大

这个时候我们就要对mmap做一次缩容操作了

很简单，申请一个更小的mmap，将老mmap拷贝过去即可

```
if d.end.chunck-d.begin.chunck < len(d.mmap)/2 && len(d.mmap) > MapSize {
    //new mmap
    mmap := make([]chunck, (d.end.chunck - d.begin.chunck + 3))

    //copy
    copy(mmap[1:len(mmap)-1], d.mmap[d.begin.chunck:d.end.chunck+1])

    //reindex
    d.begin.chunck = 1
    d.end.chunck = len(mmap) - 2

    d.mmap = mmap
```


## 说明

这个deque并不是线程(groutine)安全的,而迭代器很可能在任何一次chunck移动的过程中失效