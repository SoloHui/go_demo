/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/5
publishedTime: undefined
---

## 函数（续）

当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

在本例中，

```
x int, y int
```

被简写为

```
x, y int
```
*/
package main

import "fmt"

// 函数结构: func 名字(参数 列表) 返回值 列表
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
