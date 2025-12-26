/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/24
publishedTime: undefined
---

## 函数值

函数也是值。它们可以像其他值一样传递。

函数值可以用作函数的参数或返回值。
*/
package main

import (
	"fmt"
	"math"
)

// 相当于传入一个函数类型为 `func(float64, float64) float64` 的参数 `fn`(函数指针)
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
