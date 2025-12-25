/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/basics/2
publishedTime: undefined
---

## 导入

此代码用圆括号将导入的包分成一组，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

```javascript
import "fmt";
import "math";
```

不过使用分组导入语句要更好。

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// math.Sqrt(x) 函数返回 x 的平方根。
	fmt.Printf("现在你有了 %g 个问题。\n", math.Sqrt(7))
}
