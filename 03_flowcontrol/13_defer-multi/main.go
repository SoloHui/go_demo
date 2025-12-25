/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/flowcontrol/13
publishedTime: undefined
---

## defer 栈

推迟调用的函数调用会被压入一个栈中。

当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。

更多关于 defer 语句的信息，

请阅读[此博文](http://blog.go-zh.org/defer-panic-and-recover)。
*/
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
