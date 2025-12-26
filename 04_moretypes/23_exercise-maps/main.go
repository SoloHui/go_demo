/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/23
publishedTime: undefined
---

## 练习：映射

实现 `WordCount`。它应当返回一个映射，其中包含字符串 `s` 中每个“单词”的个数。

函数 `wc.Test` 会为此函数执行一系列测试用例，并输出成功还是失败。

你会发现 [strings.Fields](https://go-zh.org/pkg/strings/#Fields) 很有用。
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	// `strings.Fields` 会将字符串 `s` 拆分为单词切片
	words := strings.Fields(s)
	for _, word := range words {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
