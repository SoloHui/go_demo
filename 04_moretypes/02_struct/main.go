/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/2
publishedTime: undefined
---

## 结构体

一个 结构体（`struct`）就是一组 字段（field）。
*/
package main

import "fmt"

// 关键字 `type` 用于定义新的类型，`struct` 用于定义结构体。
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
