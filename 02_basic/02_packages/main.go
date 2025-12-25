/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/1
publishedTime: undefined
---

## 包

每个 Go 程序都由包构成。

程序从 `main` 包开始运行。

本程序通过导入路径 `"fmt"` 和 `"math/rand"` 来使用这两个包。

按照约定，包名与导入路径的最后一个元素一致。

例如，`"math/rand"` 包中的源码均以 `package rand` 语句开始。
*/

package main

import (
	"fmt"       // 导入 fmt 包以使用格式化输入输出功能
	"math/rand" // 导入 math/rand 包以使用随机数生成功能
)

func main() {
	// rand.Intn(n) 函数返回一个介于 0 和 n-1 之间的随机整数。
	fmt.Println("我最喜欢的数字是 ", rand.Intn(10))
}
