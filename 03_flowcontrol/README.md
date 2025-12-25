# Go 语言流程控制 (03_flowcontrol)

本目录包含 Go 语言流程控制相关的学习示例代码,主要参考自 [Go 语言之旅](https://tour.go-zh.org/)。

## 目录结构

- [01_for](#01_for) - for 循环基础
- [02_for-continued](#02_for-continued) - for 循环简化形式
- [03_for-is-gos-while](#03_for-is-gos-while) - for 作为 while 使用
- [04_forever](#04_forever) - 无限循环
- [05_if](#05_if) - if 条件判断
- [06_if-with-a-short-statement](#06_if-with-a-short-statement) - if 与简短语句
- [07_if-and-else](#07_if-and-else) - if-else 结构
- [08_exercis-loops-and-functions](#08_exercis-loops-and-functions) - 练习:循环与函数
- [09_switch](#09_switch) - switch 分支语句
- [10_switch-evaluation-order](#10_switch-evaluation-order) - switch 求值顺序
- [11_switch-with-no-condition](#11_switch-with-no-condition) - 无条件 switch
- [12_defer](#12_defer) - defer 延迟执行
- [13_defer-multi](#13_defer-multi) - defer 栈

---

## 01_for

**for 循环基础**

Go 只有一种循环结构：`for` 循环。

基本的 `for` 循环由三部分组成,用分号隔开:

```go
for 初始化语句; 条件表达式; 后置语句 {
    // 循环体
}
```

**三个组成部分:**
- **初始化语句**: 在第一次迭代前执行
- **条件表达式**:  在每次迭代前求值
- **后置语句**: 在每次迭代的结尾执行

**示例:**
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
fmt.Println(sum) // 输出: 45
```

**关键点:**
- 初始化语句中声明的变量仅在 `for` 语句的作用域中可见
- 条件表达式求值为 `false` 时,循环终止
- 与 C/Java/JavaScript 不同,Go 的 `for` 语句后面没有小括号 `()`
- 大括号 `{}` 是必须的

---

## 02_for-continued

**for 循环简化形式**

初始化语句和后置语句是**可选的**。

```go
sum := 1
for ; sum < 1000; {
    sum += sum
}
fmt.Println(sum) // 输出: 1024
```

**关键点:**
- 可以省略初始化语句和后置语句
- 分号仍然需要保留(在这种形式下)

---

## 03_for-is-gos-while

**for 是 Go 中的「while」**

当省略分号时,`for` 循环就相当于其他语言中的 `while` 循环。

```go
sum := 1
for sum < 1000 {
    sum += sum
}
fmt.Println(sum) // 输出: 1024
```

**关键点:**
- Go 中没有 `while` 关键字
- `for` 可以模拟 `while` 的功能
- 这是最常见的循环简化形式

---

## 04_forever

**无限循环**

如果省略循环条件,循环就不会结束。

```go
for {
    // 无限循环
}
```

**关键点:**
- 无限循环写法非常简洁
- 通常配合 `break` 或 `return` 语句使用
- 常用于服务器程序等需要持续运行的场景

---

## 05_if

**if 条件判断**

Go 的 `if` 语句与 `for` 循环类似,表达式外无需小括号 `()`,但大括号 `{}` 是必须的。

```go
func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}
```

**关键点:**
- 条件表达式不需要小括号
- 大括号是必须的,即使只有一条语句
- 条件表达式必须是布尔类型

---

## 06_if-with-a-short-statement

**if 与简短语句**

`if` 语句可以在条件表达式前执行一个简短语句。

```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}
```

**关键点:**
- 简短语句中声明的变量作用域**仅在 `if` 之内**(包括所有 `else` 块)
- 这种方式可以减少变量的作用域范围
- 变量 `v` 在 `if` 外部不可访问

**语法结构:**
```go
if 初始化语句; 条件表达式 {
    // if 块
}
```

---

## 07_if-and-else

**if 和 else**

在 `if` 的简短语句中声明的变量同样可以在对应的任何 `else` 块中使用。

```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // v 在此处不可用
    return lim
}
```

**关键点:**
- 在 `if` 简短语句中声明的变量在 `else` 块中也可访问
- 变量作用域覆盖整个 `if-else` 结构
- 在 `if-else` 外部无法访问该变量

**格式化输出:**
- `%g` 用于格式化浮点数,自动选择 `%e` 或 `%f` 格式

---

## 08_exercis-loops-and-functions

**练习:  循环与函数 (牛顿法求平方根)**

实现一个使用牛顿法计算平方根的函数。

**牛顿法公式:**
```go
z -= (z*z - x) / (2*z)
```

**实现:**
```go
func Sqrt(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2 * z)
        fmt.Println(z)
    }
    return z
}
```

**算法说明:**
- 从初始猜测值 `z = 1.0` 开始
- 每次迭代改进 `z` 的值
- `z*z - x` 是当前值到目标值的距离
- `2*z` 是 `z²` 的导数
- 迭代多次后可得到精确的平方根

**进阶练习:**
1. 修改循环条件,当值停止改变时退出循环
2. 尝试不同的初始猜测值(如 `x` 或 `x/2`)
3. 比较结果与 `math.Sqrt()` 的差异

**数学背景:**
- 这是[牛顿法](https://zh.wikipedia.org/wiki/%E7%89%9B%E9%A1%BF%E6%B3%95)的应用
- 通过导数来确定函数的变化速度
- 对平方根等函数非常高效

---

## 09_switch

**switch 分支语句**

`switch` 是编写一连串 `if-else` 语句的简便方法。

```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt. Println("macOS.")
case "linux":
    fmt.Println("Linux.")
default:
    fmt.Printf("%s.\n", os)
}
```

**Go 的 switch 特点:**

| 特性 | Go | C/Java/JavaScript |
|------|-----|-------------------|
| 自动 break | ✓ 自动终止 | ✗ 需要手动 `break` |
| case 类型 | 任意类型 | 通常只能是整数/字符 |
| case 值 | 可以是变量/表达式 | 通常是常量 |
| fallthrough | 需要显式使用 | 默认行为 |

**关键点:**
- Go 只运行选定的 case,自动添加 `break`
- 除非使用 `fallthrough` 语句,否则自动终止
- `case` 无需为常量,取值不限于整数
- 可以在 switch 语句中包含初始化语句

---

## 10_switch-evaluation-order

**switch 的求值顺序**

`switch` 的 `case` 语句从上到下顺次执行,直到匹配成功时停止。

```go
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
```

**关键点:**
- case 从上到下依次求值
- 匹配成功后立即停止,不再评估后续 case
- 如果 `i==0`,`case f():` 中的 `f()` 不会被调用

**注意:**
- Go 练习场中的时间总是从 **2009-11-10 23:00:00 UTC** 开始

---

## 11_switch-with-no-condition

**无条件 switch**

无条件的 `switch` 等同于 `switch true`,可以将长串 `if-else` 写得更清晰。

```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("早上好！")
case t.Hour() < 17:
    fmt. Println("下午好！")
default:
    fmt.Println("晚上好！")
}
```

**等价的 if-else 写法:**
```go
if t.Hour() < 12 {
    fmt.Println("早上好！")
} else if t.Hour() < 17 {
    fmt.Println("下午好！")
} else {
    fmt.Println("晚上好！")
}
```

**关键点:**
- `switch` 后没有表达式,等同于 `switch true`
- 每个 case 可以是任意布尔表达式
- 这种形式使代码更清晰、更易读
- 常用于替代复杂的 if-else 链

---

## 12_defer

**defer 延迟执行**

`defer` 语句会将函数推迟到外层函数返回之后执行。

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

**输出:**
```
hello
world
```

**关键点:**
- `defer` 的函数调用会推迟到外层函数返回之后执行
- 推迟调用的函数**参数会立即求值**
- 但函数调用要等到外层函数返回前才会执行

**常见用途:**
- 关闭文件:  `defer file.Close()`
- 释放锁: `defer mutex.Unlock()`
- 记录函数执行时间
- 错误处理和资源清理

---

## 13_defer-multi

**defer 栈 (后进先出)**

推迟的函数调用会被压入一个栈中,按照**后进先出**的顺序执行。

```go
func main() {
    fmt.Println("counting")
    
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("done")
}
```

**输出:**
```
counting
done
9
8
7
6
5
4
3
2
1
0
```

**执行顺序说明:**
1. 首先输出 "counting"
2. 循环中 10 个 defer 被压入栈中 (0, 1, 2, ..., 9)
3. 然后输出 "done"
4. 函数返回前,defer 栈按后进先出顺序执行 (9, 8, 7, .. ., 0)

**关键点:**
- 多个 defer 按照**后进先出** (LIFO) 的顺序执行
- defer 的参数在 defer 语句执行时就已经求值
- 这种机制类似于栈的数据结构

**更多资源:**
- [Defer, Panic, and Recover](http://blog.go-zh.org/defer-panic-and-recover) - 官方博文

---

## 流程控制对比表

| 控制结构 | Go 特点 | 与其他语言的区别 |
|---------|---------|-----------------|
| **for** | 唯一的循环结构 | 可模拟 while 和无限循环 |
| **if** | 可包含初始化语句 | 无需括号,必须有大括号 |
| **switch** | 自动 break | 不需要手动添加 break |
| **defer** | 延迟执行 | 独特的 Go 特性 |

## 运行示例

每个子目录都包含一个 `main.go` 文件,可以直接运行: 

```bash
# 运行某个示例
cd 01_for
go run main. go

# 或者直接指定路径
go run 01_for/main.go
```

## 学习顺序建议

建议按照以下顺序学习: 

1. **循环** (01-04): 掌握 `for` 循环的各种形式
2. **条件** (05-07): 学习 `if-else` 条件判断
3. **练习** (08): 通过实际问题巩固循环和函数
4. **分支** (09-11): 理解 `switch` 的使用和特性
5. **延迟执行** (12-13): 了解 `defer` 的机制和应用

## 关键概念总结

### for 循环
- Go 只有 `for` 一种循环
- 可以省略初始化和后置语句
- 省略所有部分就是无限循环

### if 语句
- 条件表达式不需要括号
- 可以包含简短语句
- 简短语句中的变量作用域限于 if-else 块

### switch 语句
- 自动 break,无需手动添加
- case 可以是任意类型和表达式
- 无条件 switch 等同于 switch true

### defer 语句
- 延迟到函数返回后执行
- 参数立即求值,调用延迟执行
- 多个 defer 按后进先出顺序执行

## 参考资源

- [Go 语言之旅 - 流程控制](https://tour.go-zh.org/flowcontrol/1) - 官方交互式教程
- [Effective Go - 控制结构](https://go.dev/doc/effective_go#control-structures) - 官方最佳实践
- [Defer, Panic, and Recover](http://blog.go-zh.org/defer-panic-and-recover) - 深入理解 defer
- [牛顿法](https://zh.wikipedia.org/wiki/%E7%89%9B%E9%A1%BF%E6%B3%95) - 数学背景

---

**注意:** 本目录的示例代码展示了 Go 语言的流程控制特性,这些特性与其他编程语言有显著差异,建议仔细体会 Go 的设计哲学。