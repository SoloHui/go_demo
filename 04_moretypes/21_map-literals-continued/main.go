/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/21
publishedTime: undefined
---

## 映射字面量（续）

若顶层类型只是一个类型名，那么你可以在字面量的元素中省略它。
*/
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
