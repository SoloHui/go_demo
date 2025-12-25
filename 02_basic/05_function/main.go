/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/4
publishedTime: undefined
---

## 函数

函数可接受零个或多个参数。

在本例中，`add` 接受两个 `int` 类型的参数。

注意类型在变量名的 **后面**。

（参考这篇关于 [Go 声明语法](http://blog.go-zh.org/gos-declaration-syntax) 的文章，了解为何使用这种类型声明的形式。）

*/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
