/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/26
publishedTime: undefined
---

## 练习：斐波纳契闭包

让我们用函数做些好玩的。

实现一个 `fibonacci` 函数，它返回一个函数（闭包），

该闭包返回一个[斐波纳契数列](https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97)

(0, 1, 1, 2, 3, 5, ...)。
*/
package main

// fibonacci 是返回一个「返回一个 int 的函数」的函数
func fibonacci() func() int {
	a, b := 0, 1
	print("-", a, b, "\n")
	return func() int {
		print("+", a, b, "\n")
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		// fmt.Println()
		f()
		/*
			-01
			+01
			+11
			+12
			+23
			+35
			+58
			+813
			+1321
			+2134
			+3455
		*/
	}
}
