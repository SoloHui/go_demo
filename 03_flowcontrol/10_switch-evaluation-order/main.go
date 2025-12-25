/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/flowcontrol/10
publishedTime: undefined
---

## switch 的求值顺序

`switch` 的 `case` 语句从上到下顺次执行，直到匹配成功时停止。

例如，

```
switch i {
case 0:
case f():
}
```

在 `i==0` 时，`f` 不会被调用。）

**注意**： Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始，

该值的意义留给读者去发现。
*/
package main

import (
	"fmt"
	"time" // 导入 time 包以使用时间相关功能
)

func main() {
	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}
}
