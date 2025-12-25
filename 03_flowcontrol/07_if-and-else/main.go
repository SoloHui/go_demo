/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/flowcontrol/7
publishedTime: undefined
---

## if 和 else

在 `if` 的简短语句中声明的变量同样可以在对应的任何 `else` 块中使用。

（在 `main` 的 `fmt.Println` 调用开始前，

两次对 `pow` 的调用均已执行并返回其各自的结果。）
*/
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// `%g` 是 `fmt.Printf` 函数的动词，用于格式化浮点数。
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
