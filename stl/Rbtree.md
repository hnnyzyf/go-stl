## 前言

最近在详细学习红黑树的实现，学习了大量网上的资料，发现对红黑树的插入，删除操作的描述，总是缺少点东西,下边的内容是阅读了大量资料并且实际测试后的记过


## 数据结构和语言

红黑树使用Golang实现

节点定义，节点其实可以定义为一个接口实现如下内容


```
type Node interface{
    GetVal() interface{}
    GetColor() bool
    Left() Node
    Right() Node
    Parent() Node
}
```

我们懒省事，使用如下节点来定义一个链表（这里体会到如果有泛型的好处了）
```
type Node struct {
    Val   int //测试使用int比较好操作
    Color int

    Left  *Node
    Right *Node

    Parent *Node
}
```

红黑树定义如下


```
type RBTree struct {
    Root *Node
}
```
我们所有的操作都是围绕着这两个数据结构来进行的，是不是很简单？

## 性质

红黑树的性质大家都很清楚，我这里只是简单的列一下：
1. 每个节点非红即黑
2. 根节点是黑的
3. 叶节点是黑的(指nil节点)
4. 如果一个节点是红的，那么它的儿子都是黑的
5. 对于任意节点，到所有叶节点的路径均包含相同数目的黑节点

网上的很多资料都只介绍了红黑树的性质，却没有告诉我们，具体怎么验证一棵树是不是红黑树，其实，我们只需要验证一棵树是不是满足这5个性质即。

##### 性质1：非红即黑
我们使用int值来存储红黑树的颜色，定义如下常量，很明显满足，所以我们不需要额外去验证

```
const (
    RED = iota
    BLACK
)
```
##### 性质2: 根节点是黑的

```
func testRoot(n *Node) bool {
    if n.Color == RED {
        return false
    }
    return true
}
```
##### 性质3：叶节点是黑的(指nil节点)
这个显而易见满足

##### 性质4：如果一个节点是红的，那么它的儿子都是黑的


```
//test RedNode
func testRedNode(n *Node) bool {
    if n.Color == BLACK {
        return true
    }

    if IsBlack(n.Left) && IsBlack(n.Right) {
        return true
    } else {
        return false
    }
}
```

##### 性质5：路径相同性
我们遍历每一个节点，计算其到所有nil节点包含的黑节点的个数

这个思路很简单：

首先，我们需要对树进行遍历，我们可以选择广度优先遍历


```
//Hook 函数便于传递对节点的操作，设计模式中有一个很经典的模式叫做Visitor 设计模式，思路就是如此

type Hook func(n *Node) bool

//do bfs for a RBtree
func (r *RBTree) BFS(hook Hook) bool {
    if r.Root == nil {
        return true
    }

    queue := []*Node{r.Root}

    idx := 0
    end := len(queue)

    for idx < end {
        n := queue[idx]

        if !hook(n) {
            return false
        }

        if n.Left != nil {
            queue = append(queue, n.Left)
            end++
        }

        if n.Right != nil {
            queue = append(queue, n.Right)
            end++
        }

        idx++
    }

    return true
}
```

然后，定义一个hook函数，计算每个节点到所有nil的路径长度

其实路径的计算本身是一个递归的过程，可以通过动态规划实现

一个节点到所有nil节的距离相等，等价于其左右节点到其对应的nil节点距离相等，当然算法可以优化的，比如加个空间缓存下每个节点是否已经检测过，检测过的就可以不检测了，咱只是测试不搞那么复杂


```
func calPath(n *Node) (int, bool) {
    if n == nil {
        return 0, true
    }

    if n.Color == BLACK {
        lc, ok1 := calPath(n.Left)
        rc, ok2 := calPath(n.Right)
        if lc == rc && ok1 && ok2 {
            return lc + 1, true
        } else {
            return -1, false
        }
    } else {
        lc, ok1 := calPath(n.Left)
        rc, ok2 := calPath(n.Right)
        if lc == rc && ok1 && ok2 {
            return lc, true
        } else {
            return -1, false
        }
    }
}
```

好了，综合以上的内容，我们就可以写出所有的test函数了



```
//性质2
if !testRoot(r.Root) {
    fmt.Println("The root is RED")
    return false
}

//性质4
if !r.BFS(testRedNode) {
    fmt.Println("The RED node does not have 2 BLACK nodes")
    return false
}

//性质5
if !r.BFS(testPath) {
    fmt.Println("The Path is not equal")
    return false
}
```

## 插入和删除

下面开始介绍插入和删除，先介绍插入，后介绍删除

我们先对节点定义进行介绍：

- n: 当前的节点，即新插入的节点
- pa: n的父亲节点
- bro: n的兄弟节点
- gp: n的祖父节点
- uc: n的叔叔节点,pa的兄弟节点,gp的子节点


### 插入
插入分为两个步骤
- 添加一个红色节点到指定位置（具体参考二叉树添加节点）
```
add操作
```

- 以该红色节点为起点，进行红黑树的平衡操作 
```
balance操作
```


所以我们可以定义如下Push函数

```
//push a new node
func (r *RBTree) Push(val int) {
    //do only once
    if r.Root == nil {
        r.init(val)
    }

    //add a new node success
    if n, ok := r.add(val); ok {
        r.Balance(n)
    }

}
```

添加红色节点这个很简单，不在复述，咱们看看如何对红黑树进行平衡，即Balance操作


```
func (r *RBTree) Balance(n *Node) {
    ......
}
```

##### 终结情况

Balance操作是一个递归操作,一个递归操作首先就要设定终结条件,终结条件有如下三种:

1. n是根节点，直接将n设置成黑色,情况1返回

```
if n == r.Root {
        n.Color = BLACK
        return
}
```

2. n的父亲节点是黑色的,情况2返回

```
if n.Parent.Color == BLACK {
        return
}
```
3. n的父亲节点(pa)是红色,叔叔节点(uc)是黑色或者nil, n,pa,gp位于一条直线


```
什么叫位于一条直线？
n == pa.Left && pa == gp.Left 或者 n == pa.Right && pa == gp.Right
```
此时，我们以pa为轴旋转,将gp节点作为pa的孩子节点（左旋右旋的定义就是这么来的）


```
如果gp变成了pa的右孩子,我们称之为右旋

如果gp变成了pa的做孩子，我们称之为左旋
```


```
if n == pa.Left && pa == gp.Left {
    r.rotate2Right(pa)
    return
}
if n == pa.Right && pa == gp.Right {
    r.rotate2Left(pa)
    return
}
```


旋转完毕后，情况3结束返回

总结一下，这三种情况下，做完操作后，整个红黑树的平衡就已经完成了，不需要再进行迭代了

##### 非终结情况

情况4：n的父亲节点(pa)是红色,叔叔节点(uc)是红色, 我们只需要将
pa和uc节点变为黑色，gp节点变为红色
以gp节点作为新插入的节点,进行递归


```
if pa.Color == RED && uc != nil && uc.Color == RED {

        //change color
        pa.Color = BLACK
        uc.Color = BLACK
        gp.Color = RED
        
        //continue Balance，set grandparent node as current node to rebalance
        r.Balance(gp)
        return
}
```

情况5：n的父亲节点(pa)是红色,叔叔节点(uc)是黑色或者nil, n,pa,gp不位于一条直线


```
什么叫不位于一条直线？
n == pa.Left && pa == gp.Right 或者 n == pa.Right && pa == gp.Left
```

此时，我们只需要以n节点为轴，将pa节点作为n的孩子节点（左旋右旋的定义就是这么来的）进行左旋右旋
此时，整个树会变成情况3

再以pa为插入节点，继续遍历即可


```
if pa.Right == n && gp.Left == pa {
    r.rotate2Left(n)
    r.Balance(pa)
    return
}

if pa.Left == n && gp.Right == pa {
    r.rotate2Right(n)
    r.Balance(pa)
    return
}
```

至此,插入操作结束


### 删除

删除操作分为如下几个步骤：

- 找到要删除的节点位置
 
 ```
 find操作
 ```
- 找到可以替换待删除节点的节点并替换之（左子树的最大值或者右子树的最小值）
 
 ```
 replace操作
 ```
- 删除用于替换的节点
 
 ```
delete操作
 ```
- 以该红色节点的兄弟节点为起点，进行红黑树的平衡操作
 
 ```
 ReBalance操作
 ```

所以我们可以定义如下pop函数
```
//pop will remove a node
func (r *RBTree) Pop(val int) {
    if r.Root == nil {
        return
    }

    //find the node which will be deleted
    if n := r.find(val); n != nil {
        //find the node which will be deleted
        del := r.replace(n)
        //exchange the value
        n.Val = del.Val
        //delete the node and get the child node
        bro, ok := r.delete(del)
        if !ok {
            r.ReBalance(bro)
        }
    }

}
```

##### 什么时候删除操作需要再平衡？

并不是所有删除节点的操作都需要再平衡的
删除的节点有如下三种情况:
- 根节点

    删除根节点表明当前的红黑树只剩一个根节点，直接删除即可

```
if n == r.Root {
    r.Root = nil
    return nil, true
}
```


- 0个child的节点
    
    如果n是红色的，删除即可，直接删除

    如果n是黑色的，需要再平衡，这里有人会问，bro为nil怎么办？
    
    其实这个是不可能的,只有红色节点才可能有nil兄弟
    
    所以，我们其实可以得到一个很好用的推论，称之为性质6:
    
    **性质6：黑色的节点必然有兄弟节点**


- 1个child的节点

    如果n是红色的，用child替换即可

    如果n是黑色的，当child是红色，将child变为黑，当child是黑色，需要再平衡

所以总结下来，需要再平衡的情况有如下两种:

    1.n是黑色并且没有孩子
    
    2.n是黑色且唯一的孩子也是黑色
 
##### 再平衡

**关键：每次平衡的入口节点为n的brother**

Reblance的过程仍然是一个递归的过程，我们先看终结条件


##### 终结情况

1. n是根节点,不需要做任何操作，情况1返回

2. n是黑色节点,n有一个红色孩子c，且pa,n,c位于一条直线上

    首先设置颜色：
    
        n的颜色设置为pa的颜色
        
        pa的颜色设置为黑色
        
        c 的颜色设置为黑色

    然后,我们只需要以n为轴，将pa作为n的孩子节点进行左旋右旋
    
    情况2返回
    

```
if pa.Right == n && IsRed(n.Right) {

        n.Right.Color = BLACK
        n.Color = pa.Color
        pa.Color = BLACK

        r.rotate2Left(n)
        return

}

if pa.Left == n && IsRed(n.Left) {
            
        n.Left.Color = BLACK
        n.Color = pa.Color
        pa.Color = BLACK

        r.rotate2Right(n)
        return
}
```

    
3. n是黑色节点，且两个孩子为黑色或者nil,pa是红色的话

    设置颜色：
    
        n的颜色设置为红色
    
        pa的颜色设置为黑色

    情况3返回



##### 非终结情况

4. n是红色节点

    设置颜色:
        
        n的颜色设置为黑色
        pa的颜色设置为红色
        
    以n为轴，将pa作为n的孩子节点进行左选右选
    
    以pa的左或右节点再平衡

```
n.Color = BLACK
pa.Color = RED
        
if pa.Left == n {
    r.rotate2Right(n)
    r.ReBalance(pa.Left)
    return
} else {
    r.rotate2Left(n)
    r.ReBalance(pa.Right)
    return
}
```


5. n是黑色节点，且两个孩子为黑色或者nil,pa是 黑色的话

    设置颜色:
        
        n的颜色设置为红色
    
    以pa的兄弟节点再平衡
    
6. n是黑色节点，有一个红色的孩子c，且该孩子pa,n,c不在一条直线上

    首先设置颜色：
    
        n的颜色设置为红色
        
        c 的颜色设置为黑色

    然后,我们只需要以c节点进行左旋或者右旋
    
    最后以c节点再平衡
    

```
if pa.Right == n && IsRed(n.Left) {
    c := n.Left

    n.Color = RED
    c.Color = BLACK

    r.rotate2Right(c)

    r.ReBalance(c)
    return

}

if pa.Left == n && IsRed(n.Right) {
    //child 
    c := n.Right

    //set color
    n.Color = RED
    c.Color = BLACK

    //rotate
    r.rotate2Left(c)

    //rebalance
    r.ReBalance(c)
    return
}
```



##### 左旋和右旋

左旋:将当前节点n的父亲节点作为自己的左子树

右旋:将当前节点n的父亲节点作为自己的右子树

```
func (r *RBTree) rotate2Left(n *Node) {
    pa := n.Parent
    gp := pa.Parent

    pa.Right = n.Left
    if n.Left != nil {
        n.Left.Parent = pa
    }

    //set pa as left child
    n.Left = pa
    pa.Parent = n

    if gp == nil {
        r.Root = n
    } else if gp.Left == pa {
        gp.Left = n
    } else {
        gp.Right = n
    }
    n.Parent = gp
}

func (r *RBTree) rotate2Right(n *Node) {
    pa := n.Parent
    gp := pa.Parent

    pa.Left = n.Right
    if n.Right != nil {
        n.Right.Parent = pa
    }

    //set pa as right child
    n.Right = pa
    pa.Parent = n
    if gp == nil {
        r.Root = n
    } else if gp.Left == pa {
        gp.Left = n
    } else {
        gp.Right = n
    }
    n.Parent = gp

}
```

## 为什么我们需要红黑树


算法 | 最坏查找 | 最坏插入 | 平均查找 | 平均插入 | 是否有序
---  |  --- | --- | --- | --- | ---|
顺序查找 |N |   N|  N/2 |N| 否
二分查找 |lgN|  N|  lgN|    N/2 |是
二叉树查找（BST） | N| N|  1.39lgN|    1.39lgN|    是
红黑树查找（RBT） |2lgN|   2lgN|   1.001lgN|   1.001lgN|   是



### 无锁红黑树的实现
https://xuezhaokun.github.io/150-algorithm/



























