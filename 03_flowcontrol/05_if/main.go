/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/flowcontrol/5
publishedTime: undefined
---

## if 判断

Go 的 `if` 语句与 `for` 循环类似，表达式外无需小括号 `( )`，

而大括号 `{ }` 则是必须的。
*/
package main

import (
	"fmt"
	"math"
)

// sqrt 计算并返回 x 的平方根，当 x 为负数时返回虚数部分。
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// 'math.Sqrt' 函数返回一个浮点数的平方根。
	// 'fmt.Sprint' 函数将其参数格式化并返回其字符串表示形式。
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
