# 前言
KMP是字符串匹配中一个非常经典的算法,对于后序的Trie,AC自动的学习均有很大的帮助，算法库中直接提供了KMP算法的实现


## 问题

KMP解决的问题为:判断一个字符串中是否有匹配的模式


## 算法简介

KMP的算法核心是计算Next数组，next数组中记录的是公共子串的长度，也就是从起始位置的偏移量

举例:

比如对于一个模式"ABCDABD",待匹配的字符串为"BBC ABCDAB ABCDABCDABDE"

#### next数组的意义
则其Next数组计算方法为:

子串 | 前缀 | 后缀 | 公共子串长度
---|---|---|---
A|{}|{}| 0 
AB|{A}|{B}|0
ABC|{A,AB}|{BC,C}|0
ABCD|{A,AB,ABC}|{BCD,CD,D}|0
ABCDA|{A,AB,ABC,ABCD}|{BCDA,CDA,DA,A}|1
ABCDAB|{A,AB,ABC,ABCD,ABCDA}|{BCDAB,CDAB,DAB,AB,B}|2
ABCDABD|{A,AB,ABC,ABCD,ABCDA,ABCDAB}|{BCDABD,CDABD,DABD,ABD,BD,D}|0


则其next数组为
A|B|C|D|A|B|D
---|---|---|---|---|---|---
0|0|0|0|1|2|0


#### 匹配失败
匹配字符串"BBC ABCDAB ABCDABCDABDE"，当我们匹配到ABCDAB的时候,因为待匹配字符串的下一个字符为空格，而模式的匹配字符串为D，**匹配失败**

如下图所示

```
    BBC ABCDAB ABCDABCDABDE
        ABCDABD
```

我们不会从模式头开始重新一次匹配ABCDABD,而是通过next数组来确定从pattern的哪个位置重新开始匹配,计算公式如下

next数组中已经匹配位数为2,那我们就从第三位(下标为2)的位置重新开始匹配

```
    BBC ABCDAB ABCDABCDABDE
            ABCDABD
```

直到找到对应的子串或者匹配结束


#### next数组的计算方法

我们不可能真的把每一个模式的子串的前后缀求出并找出最长公共子串,这样KMP的算法就没有意义,其实通过观察，我们可以认为KMP其实每次就是找 **当前位置向前n位和从起始位置向后n位一样的最长子串**,这样理解起来比较绕，让我们给出图来形象的说明

0|1|2|3|4|5|6
---|---|---|---|---|---|---
A | B |C |D|A|B|D
0|0|0|0|1|2|0

当i=0，很明显Next[0] = 0


当i=5的时候,Next[4] = 1
这里的意义就是从起始位置开始p[0:1] 和 p [4:5] 是同一个字符串,
如果p[5] = p[1],那么的子串长度p[0:2]或者p[4:6]的长度，很明显,满足
所以Next[5] = Next[4] +1

当i=6的时候，Next[5] = 2
很明显,p[0:2]的下个字符和p[6] 不相等,我们查看p[2]的最长子串,发现为0,
且p[0] 也不等于p[6],则Next[6] = 0


```
//很明显Next[0] = 0,所以我们从第二位开始计算，也就是i=1开始
for i := 1; i < len(m.p); i++ {
	//offset记录的是上个字符的最长子串的长度
        //该子串为从起始位置开始往后查offset位的子串
	offset := m.next[i-1]

        //找到下一个位置相同的公共子串
	for offset > 0 && m.p[i] != m.p[offset] {
                //寻找公共子串中的最长公共子串
		offset = m.next[offset-1]
	}

	//找到,offset加1
	if m.p[i] == m.p[offset] {
		offset++
	}

	m.next[i] = offset
}
```