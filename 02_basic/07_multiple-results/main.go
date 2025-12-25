/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/6
publishedTime: undefined
---

## 多重返回值

Go 的函数可以返回任意数量的返回值。

`swap` 函数接受两个字符串参数，并返回它们的值互换后的结果。

```go
func swap(x, y string) (string, string) {
```
*/
package main

import "fmt"

// swap 函数返回两个字符串值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// 符号 `:=` 用于声明并初始化变量。
	// 这里声明了两个变量 a 和 b 来接收 swap 函数的两个返回值。
	a, b := swap("hello", "world")
	fmt.Printf("a=%s b=%s\n", a, b)

	// 符号 `=` 用于给已经声明的变量赋值。
	a, b = swap("goodbye", "world")
	fmt.Printf("a=%s b=%s\n", a, b)
}
