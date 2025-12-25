# Go 语言基础 (02_basic)

本目录包含 Go 语言基础知识的学习示例代码,主要参考自 [Go 语言之旅](https://tour.go-zh.org/)。

## 目录结构

- [01_print](#01_print) - 打印输出基础
- [02_packages](#02_packages) - 包的概念
- [03_imports](#03_imports) - 导入语句
- [04_exported-names](#04_exported-names) - 导出名规则
- [05_function](#05_function) - 函数定义
- [06_functions-continued](#06_functions-continued) - 函数参数简写
- [07_multiple-results](#07_multiple-results) - 多重返回值
- [08_named-results](#08_named-results) - 命名返回值
- [09_variables](#09_variables) - 变量声明
- [10_variables-with-initializers](#10_variables-with-initializers) - 带初始化的变量
- [11_short-variable-declarations](#11_short-variable-declarations) - 短变量声明
- [12_basic-types](#12_basic-types) - 基本类型
- [13_zero](#13_zero) - 零值
- [14_type-conversions](#14_type-conversions) - 类型转换
- [15_type-inference](#15_type-inference) - 类型推断
- [16_constants](#16_constants) - 常量
- [17_numeric-constants](#17_numeric-constants) - 数值常量

---

## 01_print

**打印输出基础**

演示 Go 语言中的基本输出功能,使用 `fmt` 包进行格式化输出。

```go
fmt.Println("Hello, welcome to Go programming!")  // 打印并换行
fmt.Print("文本")                                   // 打印不换行
fmt.Printf("格式化数字: %d\n", 42)                   // 格式化输出
```

**关键点:**
- `fmt.Println` - 打印后自动换行
- `fmt.Print` - 打印不换行
- `fmt.Printf` - 格式化输出

---

## 02_packages

**包的概念**

每个 Go 程序都由包构成,程序从 `main` 包开始运行。

```go
package main

import (
    "fmt"
    "math/rand"
)
```

**关键点:**
- 每个 Go 程序都由包构成
- 程序从 `main` 包开始运行
- 包名与导入路径的最后一个元素一致,如 `"math/rand"` 包中的代码以 `package rand` 开始

---

## 03_imports

**导入语句**

演示如何导入多个包,推荐使用分组导入的形式。

**推荐写法 (分组导入):**
```go
import (
    "fmt"
    "math"
)
```

**也可以这样写 (但不推荐):**
```go
import "fmt"
import "math"
```

---

## 04_exported-names

**导出名规则**

在 Go 中,以大写字母开头的名字是已导出的(可被外部访问)。

```go
math.Pi   // ✓ 正确 - Pi 首字母大写,是导出的
math.pi   // ✗ 错误 - pi 首字母小写,未导出
```

**关键点:**
- 大写字母开头 → 已导出 (public)
- 小写字母开头 → 未导出 (private)
- 导入包时,只能引用已导出的名字

---

## 05_function

**函数定义**

函数可以接受零个或多个参数,类型在变量名**后面**。

```go
func add(x int, y int) int {
    return x + y
}
```

**函数结构:**
```
func 函数名(参数名 参数类型) 返回值类型 {
    // 函数体
}
```

---

## 06_functions-continued

**函数参数简写**

当连续的参数类型相同时,除最后一个外都可以省略类型。

```go
// 完整写法
func add(x int, y int) int

// 简写形式
func add(x, y int) int
```

---

## 07_multiple-results

**多重返回值**

Go 的函数可以返回任意数量的返回值。

```go
func swap(x, y string) (string, string) {
    return y, x
}

// 使用
a, b := swap("hello", "world")
```

**关键点:**
- `:=` 用于声明并初始化变量
- `=` 用于给已声明的变量赋值

---

## 08_named-results

**命名返回值**

Go 的返回值可以被命名,它们会被视作定义在函数顶部的变量。

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // 裸返回
}
```

**关键点:**
- 命名返回值可作为文档使用
- 没有参数的 `return` 语句返回已命名的返回值
- 裸返回仅适用于短函数,长函数中会影响可读性

---

## 09_variables

**变量声明**

`var` 语句用于声明变量,类型在变量名后面。

```go
var i int
var c, python, java bool
```

**关键点:**
- `var` 语句可以出现在包或函数层级
- 可以一次声明多个变量
- 类型在最后

---

## 10_variables-with-initializers

**带初始化的变量声明**

变量声明时可以包含初始值,如果提供了初始值,可以省略类型。

```go
var i, j int = 1, 2
var c, python, java = true, false, "no!"
```

---

## 11_short-variable-declarations

**短变量声明**

在函数中,可以使用 `:=` 简短赋值语句代替 `var` 声明。

```go
func main() {
    k := 3
    c, python, java := true, false, "no!"
}
```

**关键点:**
- `:=` 只能在函数内使用
- 函数外的语句必须以关键字开始 (`var`、`func` 等)
- `:=` 等价于 `var` 声明与初始化的结合

---

## 12_basic-types

**基本类型**

Go 的基本类型包括: 

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte     // uint8 的别名
rune     // int32 的别名,表示一个 Unicode 码位

float32 float64

complex64 complex128
```

**关键点:**
- `int`、`uint` 和 `uintptr` 在 32 位系统上为 32 位,64 位系统上为 64 位
- 需要整数值时应使用 `int`,除非有特殊理由使用固定大小或无符号整数类型
- 变量声明可以「分组」成代码块

**示例:**
```go
var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

---

## 13_zero

**零值**

没有明确初始化的变量会被赋予对应类型的零值。

| 类型 | 零值 |
|------|------|
| 数值类型 | `0` |
| 布尔类型 | `false` |
| 字符串 | `""` (空字符串) |

```go
var i int       // 0
var f float64   // 0
var b bool      // false
var s string    // ""
```

**格式化输出:**
- `%v` - 以默认格式打印值
- `%q` - 输出带双引号的字符串

---

## 14_type-conversions

**类型转换**

表达式 `T(v)` 将值 `v` 转换为类型 `T`。

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// 简短形式
i := 42
f := float64(i)
u := uint(f)
```

**关键点:**
- Go 在不同类型间赋值时需要**显式转换**
- 与 C 语言不同,Go 不支持隐式类型转换

---

## 15_type-inference

**类型推断**

声明变量但不指定类型时,变量类型由右值推断。

```go
var i int
j := i  // j 也是 int

// 数值常量的类型推断
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

**关键点:**
- 使用 `:=` 或 `var =` 时可省略类型
- 右值为确定类型时,新变量类型与其相同
- 右值为数值常量时,类型取决于常量精度

---

## 16_constants

**常量**

常量使用 `const` 关键字声明。

```go
const Pi float32 = 3.14
const World = "世界"
const Truth = true
```

**关键点:**
- 常量可以是字符、字符串、布尔值或数值
- 常量不能用 `:=` 语法声明
- 常量声明与变量类似

---

## 17_numeric-constants

**数值常量**

数值常量是高精度的**值**,未指定类型的常量由上下文决定其类型。

```go
const (
    Big   = 1 << 100  // 非常大的数
    Small = Big >> 99 // 2
)

func needInt(x int) int
func needFloat(x float64) float64
```

**关键点:**
- 数值常量可以表示非常大的数值
- 类型由使用上下文决定
- `int` 类型可存储最大 64 位整数(根据平台不同可能更小)

---

## 运行示例

每个子目录都包含一个 `main.go` 文件,可以直接运行: 

```bash
# 运行某个示例
cd 01_print
go run main.go

# 或者直接指定路径
go run 01_print/main.go
```

## 学习顺序建议

建议按照目录编号顺序学习:

1. **输入输出** (01-03): 了解基本的包、导入和输出
2. **函数** (04-08): 掌握函数的定义和使用
3. **变量** (09-11): 学习变量的声明方式
4. **类型** (12-17): 深入理解类型系统

## 参考资源

- [Go 语言之旅](https://tour.go-zh.org/) - 官方交互式教程
- [Go 语言规范](https://go.dev/ref/spec) - 官方语言规范
- [Go 声明语法](http://blog.go-zh.org/gos-declaration-syntax) - 了解 Go 的声明语法设计

---

**注意:** 本目录的示例代码主要用于学习目的,展示了 Go 语言的基础语法和核心概念。