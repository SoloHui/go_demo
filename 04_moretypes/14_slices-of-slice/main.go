/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/moretypes/14
publishedTime: undefined
---

## 切片的切片
切片可以包含任何类型，当然也包括其他切片。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字棋（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
