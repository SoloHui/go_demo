/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/12
publishedTime: undefined
---

## nil 切片
切片的零值是`nil`。

`nil`切片的长度和容量为 0 且没有底层数组。
*/
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
