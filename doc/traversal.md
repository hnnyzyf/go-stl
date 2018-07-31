## 遍历

这一篇主要介绍一下遍历,分别介绍常用的二叉树的遍历和树的遍历


### 二叉树的遍历

#### 遍历的三种分类
二叉树的遍历是非常常用的遍历方法，主要包括如下三种:

    - 前序遍历
    - 中序遍历
    - 后序遍历

三种遍历实际上均是深度优先遍历的特例,而深度优先遍历需要记录上一个状态并且在当前状态结束后，返回上一个状态
简单的说，即先进先出，所以三种遍历均需要提供一个stack来记录状态



#### 遍历的两种写法
遍历的写法分为非递归和递归的写法
    
    - 递归写法
    结构简单易于理解，使用函数栈作为状态的记录,但是递归常常会造成栈溢出(stackoverflow)
    - 非递归写法
    略有点复杂，需要使用者自定义栈结构，但是不会有栈溢出的可能性
    
    
我们后续使用非递归写法来实现遍历


#### 二叉树的定义
一个二叉树可以为如下形式

                A
            B       C

### 前序遍历

前序遍历指的如下遍历路径 

    根节点(A) -> 左子树 (B) ->右子树(C)
    根节点(A) -> 右子树 (C) ->左子树(B)

很简单的给出前序遍历的非递归伪代码，需要一个栈来存储已经遍历过得左子树（或右子树）

    新建栈S
    S.Push(root)
    while S不为空
        curr = S.Pop()
        输出
        if curr.Right存在
            S.Push(curr.Right) //保证左子树的遍历在右子树之前
        if curr.Left存在
            S.Push(curr.Left)
      

### 中序遍历

中序遍历指的如下遍历路径 

    左子树(B) -> 根节点 (A) ->右子树(C) 或者
    右子树(C) -> 根节点 (A -> 左子树(B)


很简单的给出中序遍历的非递归伪代码

    新建栈S
    S.Push(root)
    while S不为空
        curr = S.Pop()
        while curr不为空
            S.Push(curr) //保证先遍历完左子树
            curr = curr.Left

        n = S.Pop()
        输出
        curr = n.Right //当前节点的左子树遍历玩且当前节点也被遍历后，才遍历右子树



### 后序遍历

后序遍历指的如下遍历路径 

    右子树(C) ->左子树(B) -> 根节点 (A) 或者
    左子树(B) ->右子树(C) -> 根节点


很简单的给出后序序遍历的非递归伪代码

1.单栈后续遍历
需要记录当前节点是否被遍历过,如果没有被遍历过，则直接遍历有节点，如果已经被遍历过，则返回上一个状态，空间复杂度是o(h),因为只需要一个栈，但需要记录节点的状态,可以通过在节点上添加flag，或者开辟一个新的map来存储状态



2.使用双栈看起来更加简单明了，但是空间复杂度是o(n)，因为Output栈要存储所有的节点

    新建栈S
    新建栈OutPut
    S.Push(root)
    while S不为空
        curr = S.Pop()
        OutPut.Push(curr)
        if curr.Left存在
            S.Push(curr.Left)
        if curr.Right存在
            S.Push(curr.Right)

    while OutPut不为空
        n = OutPut.Pop()
        输出n


### 树的遍历

树本身就是有向图的一个特例，所以下边对树的遍历，讨论两种方法:

    -广度优先遍历(BFS)
    -深度优先便利(DFS)

#### 如下为一个树

                        A
            B           C           D
        E       F       G           H


#### 广度优先遍历

广度优先指如下过程:

            A -> B -> C -> D -> E -> F -> G -> H


因为 树的孩子不向二叉树会区分左右,故孩子的顺序可以随意确定

下边给出广度优先的遍历非递归遍历伪代码

    新建队列Q
    Q.Push(root)
    while Q不为空
        curr = Q.Pop()
        输出curr
        for child:=curr.Children()
            Q.Push(child)



#### 深度优先遍历
 
##### 前序深度遍历
前序深度遍历指如下过程:

             A - > B  -> E -> F -> C -> G -> D -> H


因为 树的孩子不区分顺序,故孩子的顺序可以随意确定，类似于二叉树的前序遍历

下边给出深度优先的遍历非递归遍历伪代码

    新建栈S
    S.Push(root)
    while S不为空
        curr = S.Pop()
        输出n
        for child:=curr.children()
            S.Push(child)


###后序深度遍历

            E->F->B->G->C->H-D->A

使用双栈看起来更加简单明了,但是空间复杂明显上升为O(n)

    新建栈S
    新建栈OutPut
    S.Push(root)
    while S不为空
        curr = S.Pop()
        OutPut.Push(curr)
        for child:=curr.children()
            S.Push(child)

    while OutPut不为空
        n = OutPut.Pop()
        输出n
