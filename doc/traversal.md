## 遍历

这一篇主要介绍一下遍历,分别介绍常用的二叉树的遍历和树的遍历


### 二叉树的遍历

二叉树的遍历是非常常用的遍历方法，主要包括如下三种,下边给出伪代码

    - 前序遍历
    - 中序遍历
    - 后序遍历

一个二叉树可以为如下形式

                A
            B       C

#### 前序遍历


前序遍历指的如下遍历路径 

    根节点(A) -> 左子树 (B) ->右子树(C)
    根节点(A) -> 右子树 (C) ->左子树(B)


很简单的给出前序遍历的非递归伪代码，需要一个栈来存储已经遍历过得左子树

    新建栈S
    S.Push(root)
    while S不为空
        n = S.Pop()
        输出n
        S.Push(n.Right)
        S.Push(n.Left)
      

#### 中序遍历


前序遍历指的如下遍历路径 

    左子树(B) -> 根节点 (A) ->右子树(C) 或者
    右子树(C) -> 根节点 (A -> 左子树(B)


很简单的给出后序序遍历的非递归伪代码，需要栈来存储

    新建栈S
    curr = root
    while S不为空 或者 curr 不为空
        while curr不为空
            S.Push(curr)
            curr = curr.Left

        if S不为空
            curr = S.Pop()
            输出curr
            curr = curr.Right



#### 后序遍历

前序遍历指的如下遍历路径 

    右子树(C) ->左子树(B) -> 根节点 (A) 或者
    左子树(B) ->右子树(C) -> 根节点

很简单的给出后序序遍历的非递归伪代码，使用双栈看起来更加简单明了

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


### 总结




### 树的遍历

树本身就是有向图的一个特例，所以下边对树的遍历，讨论两种方法:

    -深度优先便利(DFS)
    -广度优先遍历(NFS)


#### 广度优先遍历

                    A
            B       C       D
        E       F       G       H
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
                    A
            B           C          D
        E       F       G          H


##### 前序深度遍历
前序深度遍历指如下过程:

             A - > B  -> E -> F -> C -> G -> D -> H


因为 树的孩子不向二叉树会区分左右,故孩子的顺序可以随意确定，类似于二叉树的前序遍历

下边给出深度优先的遍历非递归遍历伪代码

    新建栈S
    S.Push(root)
    while S不为空
        n = S.Pop()
        输出n
        for child:=S.children()
            S.Push(child)


##### 中序或后续深度遍历

中序深度遍历指如下过程:
E->B->F->G->C->H-D->A