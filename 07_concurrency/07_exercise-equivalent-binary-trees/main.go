/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/concurrency/7
publishedTime: undefined
---

## 练习：等价二叉查找树

不同二叉树的叶节点上可以保存相同的值序列。

例如，以下两个二叉树都保存了序列 \`1，1，2，3，5，8，13\`。

![](https://tour.go-zh.org/static/img/tree.png)

在大多数语言中，检查两个二叉树是否保存了相同序列的函数都相当复杂。

我们将使用 Go 的并发和信道来编写一个简单的解法。

本例使用了 `tree` 包，它定义了类型：

```go
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```

---
title: Go 语言之旅
url: https://tour.go-zh.org/concurrency/8
publishedTime: undefined
---

## 练习：等价二叉查找树

**1.** 实现 `Walk` 函数。

**2.** 测试 `Walk` 函数。

函数 `tree.New(k)` 用于构造一个随机结构的已排序二叉查找树，

它保存了值 `k`, `2k`, `3k`, ..., `10k`。

创建一个新的信道 `ch` 并且对其进行步进：

```
go Walk(tree.New(1), ch)
```

然后从信道中读取并打印 10 个值。应当是数字 1, 2, 3, ..., 10.

**3.** 用 `Walk` 实现 `Same` 函数来检测 `t1` 和 `t2` 是否存储了相同的值。

**4.** 测试 `Same` 函数。

`Same(tree.New(1), tree.New(1))` 应当返回 `true`，

而 `Same(tree.New(1), tree.New(2))` 应当返回 `false`。

`Tree` 的文档可在[这里](https://godoc.org/golang.org/x/tour/tree#Tree)找到。

*/

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 遍历树 t，并树中所有的值发送到信道 ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	fmt.Printf("%d\n", t.Value)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 判断 t1 和 t2 是否包含相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}

	_, ok := <-ch2
	return !ok
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)

	println(Same(t1, t2)) // 应该输出 true
	println(Same(t1, t3)) // 应该输出 false
}
