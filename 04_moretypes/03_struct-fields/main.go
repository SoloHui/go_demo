/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/3
publishedTime: undefined
---

## 结构体字段

结构体字段可以通过 `.` 访问。
*/
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	v.Y = 5
	fmt.Println(v.X)
}
