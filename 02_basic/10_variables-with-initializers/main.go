/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/9
publishedTime: undefined
---

## 变量的初始化

变量声明可以包含初始值，每个变量对应一个。

如果提供了初始值，则类型可以省略；变量会从初始值中推断出类型。
*/
package main

import "fmt"

var i, j int = 1, 2

func main() {
	// 初始值推断类型
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
