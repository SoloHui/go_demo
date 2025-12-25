/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/flowcontrol/6
publishedTime: undefined
---

## if 和简短语句

和 `for` 一样，`if` 语句可以在条件表达式前执行一个简短语句。

该语句声明的变量作用域仅在 `if` 之内。

（在最后的 `return` 语句处使用 `v` 看看。）
*/
package main

import (
	"fmt"
	"math"
)

// pow 函数返回 x 的 n 次幂；如果结果小于 lim，则返回结果，否则返回 lim。
func pow(x, n, lim float64) float64 {
	// `math.Pow` 函数返回 x 的 n 次幂。
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
