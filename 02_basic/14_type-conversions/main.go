/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/13
publishedTime: undefined
---

## 类型转换

表达式 `T(v)` 将值 `v` 转换为类型 `T`。

一些数值类型的转换：

```javascript
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

或者，更加简短的形式：

```
i := 42
f := float64(i)
u := uint(f)
```

与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。试着移除例子中的 `float64` 或 `uint` 的类型转换，看看会发生什么。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// cannot use (x * x + y * y) (value of type int) as float64 value in argument to math.Sqrt
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// cannot use (f) (variable of type float64) as uint value in variable declaration
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
