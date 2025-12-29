/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/methods/9
publishedTime: undefined
---

## 接口
**接口类型**的定义为一组方法签名。

接口类型的变量可以持有任何实现了这些方法的值。

**注意:** 示例代码的第 22 行存在一个错误。

由于 Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser。
*/
package main

import (
	"fmt"
	"math"
)

// `interface` 类型 `Abser` 包含一个 `Abs` 方法签名。
// 任何实现了该方法的类型都实现了该接口。
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
