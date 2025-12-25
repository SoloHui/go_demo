/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/12
publishedTime: undefined
---

## 零值

没有明确初始化的变量声明会被赋予对应类型的 **零值**。

零值是：

- 数值类型为 `0`，
- 布尔类型为 `false`，
- 字符串为 `""`（空字符串）。
*/
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	// `%v` 以默认格式打印值
	// `%q` 输出带双引号的字符串
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
