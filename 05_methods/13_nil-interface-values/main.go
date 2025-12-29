/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/methods/13
publishedTime: undefined
---

## nil 接口值

nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误，

因为接口的元组内并未包含能够指明该调用哪个 **具体** 方法的类型。
*/
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// nil dereference in dynamic method call
	i.M()
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4992b9]
	*/
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
