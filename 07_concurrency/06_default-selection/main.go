/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/concurrency/6
publishedTime: undefined
---

## 默认选择

当 `select` 中的其它分支都没有准备好时，`default` 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 `default` 分支：

```
select {
case i := <-c:

	// 使用 i

default:

	    // 从 c 中接收会阻塞时执行
	}

```
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个定时器
	// tick 每 100 毫秒触发一次
	// boom 在 500 毫秒后触发一次
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		// 监听 tick 和 boom 的触发事件
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
