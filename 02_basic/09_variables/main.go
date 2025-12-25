/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/8
publishedTime: undefined
---

## 变量

`var` 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。

如例中所示，`var` 语句可以出现在包或函数的层级。
*/
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
