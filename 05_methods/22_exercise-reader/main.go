/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/methods/22
publishedTime: undefined
---

## 练习：Reader

实现一个 `Reader`，它产生 ASCII 字符 'A'的无限流。
*/
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 为 MyReader 添加一个 Read([]byte) (int, error) 方法。
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
