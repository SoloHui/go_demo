# Go 语言更多类型 (04_moretypes)

本目录包含 Go 语言中更多类型的学习示例代码,包括指针、结构体、数组、切片、映射和函数等高级特性。主要参考自 [Go 语言之旅](https://tour.go-zh.org/)。

## 目录结构

### 指针与结构体
- [01_pointers](#01_pointers) - 指针
- [02_struct](#02_struct) - 结构体
- [03_struct-fields](#03_struct-fields) - 结构体字段
- [04_struct-pointers](#04_struct-pointers) - 结构体指针
- [05_struct-literals](#05_struct-literals) - 结构体字面量

### 数组与切片
- [06_array](#06_array) - 数组
- [07_slices](#07_slices) - 切片
- [08_slices-pointers](#08_slices-pointers) - 切片类似数组的引用
- [09_slice-literals](#09_slice-literals) - 切片字面量
- [10_slice-bounds](#10_slice-bounds) - 切片的默认行为
- [11_slice-len-cap](#11_slice-len-cap) - 切片的长度与容量
- [12_nil-slices](#12_nil-slices) - nil 切片
- [13_making-slices](#13_making-slices) - 用 make 创建切片
- [14_slices-of-slice](#14_slices-of-slice) - 切片的切片
- [15_append](#15_append) - 向切片追加元素
- [16_range](#16_range) - range 遍历
- [17_range-continued](#17_range-continued) - range 遍历(续)
- [18_exercise-slices](#18_exercise-slices) - 练习:切片

### 映射 (Map)
- [19_maps](#19_maps) - map 映射
- [20_map-literals](#20_map-literals) - 映射字面量
- [21_map-literals-continued](#21_map-literals-continued) - 映射字面量(续)
- [22_mutating-maps](#22_mutating-maps) - 修改映射
- [23_exercise-maps](#23_exercise-maps) - 练习:映射

### 函数值与闭包
- [24_function-values](#24_function-values) - 函数值
- [25_function-closures](#25_function-closures) - 函数闭包
- [26_exercise-fibonacci-closure](#26_exercise-fibonacci-closure) - 练习:斐波纳契闭包

---

## 指针与结构体

### 01_pointers

**指针**

Go 拥有指针,指针保存了值的内存地址。

**指针类型:**
- 类型 `*T` 是指向 `T` 类型值的指针
- 零值为 `nil`

**指针操作符:**

| 操作符 | 说明 | 示例 |
|-------|------|------|
| `&` | 取地址操作符,生成指向操作数的指针 | `p = &i` |
| `*` | 解引用操作符,表示指针指向的底层值 | `*p = 21` |

**示例:**
```go
i, j := 42, 2701

p := &i         // 指向 i
fmt.Println(*p) // 通过指针读取 i 的值:  42
*p = 21         // 通过指针设置 i 的值
fmt.Println(i)  // 查看 i 的值: 21

p = &j         // 指向 j
*p = *p / 37   // 通过指针对 j 进行除法运算
fmt.Println(j) // 查看 j 的值: 73
```

**关键点:**
- 与 C 不同,Go **没有指针运算**
- `&` 获取变量地址
- `*` 访问指针指向的值(解引用)

---

### 02_struct

**结构体**

一个结构体 (`struct`) 就是一组字段 (field)。

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})  // 输出: {1 2}
}
```

**关键点:**
- 使用 `type` 关键字定义新类型
- 使用 `struct` 关键字定义结构体
- 结构体是字段的集合

---

### 03_struct-fields

**结构体字段**

结构体字段使用 `.` 访问。

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    v.Y = 5
    fmt.Println(v. X)  // 输出: 4
}
```

**关键点:**
- 使用 `.` 操作符访问结构体字段
- 可以读取和修改字段的值

---

### 04_struct-pointers

**结构体指针**

结构体字段可通过结构体指针来访问。

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v
    
    // 显式解引用
    (*p).X = 210
    fmt.Println(v)  // {210 2}
    
    // 隐式解引用 (推荐)
    p.X = 1e9
    fmt.Println(v)  // {1000000000 2}
}
```

**关键点:**
- 可以通过 `(*p).X` 访问结构体字段 (显式解引用)
- Go 允许隐式解引用,直接写 `p.X` 更简洁
- 隐式解引用是 Go 的语法糖

---

### 05_struct-literals

**结构体字面量**

结构体字面量通过列出字段的值来新分配一个结构体。

```go
type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
    v2 = Vertex{X:  1}  // Y: 0 被隐式地赋予零值
    v3 = Vertex{}      // X:0 Y:0
    p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体(指针)
)
```

**初始化方式:**

| 方式 | 说明 | 示例 |
|------|------|------|
| 按顺序列出 | 按字段定义顺序赋值 | `Vertex{1, 2}` |
| 使用 `Name:` 语法 | 可以只列出部分字段 | `Vertex{X:  1}` |
| 空结构体 | 所有字段为零值 | `Vertex{}` |
| `&` 前缀 | 返回指向结构体的指针 | `&Vertex{1, 2}` |

---

## 数组与切片

### 06_array

**数组**

类型 `[n]T` 表示一个数组,它拥有 `n` 个类型为 `T` 的值。

```go
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1])  // Hello World
fmt.Println(a)           // [Hello World]

primes := [6]int{2, 3, 5, 7, 11, 13}
fmt. Println(primes)  // [2 3 5 7 11 13]
```

**关键点:**
- 数组的长度是其类型的一部分
- 数组**不能改变大小**
- 声明语法: `[长度]类型`
- 这看起来是个限制,但 Go 有更方便的切片

---

### 07_slices

**切片**

切片为数组元素提供动态大小的、灵活的视角。在实践中,切片比数组更常用。

**语法:**
```go
a[low:high]  // 半闭半开区间:  包括 low,排除 high
```

**示例:**
```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]
fmt.Println(s)  // [3 5 7]
```

**关键点:**
- 类型 `[]T` 表示一个元素类型为 `T` 的切片
- 切片通过两个下标界定:  下界和上界
- 半闭半开区间:  包括第一个元素,排除最后一个元素
- 切片是对数组的灵活视图

---

### 08_slices-pointers

**切片类似数组的引用**

切片并不存储任何数据,它只是描述了底层数组中的一段。

```go
names := [4]string{"John", "Paul", "George", "Ringo"}
fmt.Println(names)  // [John Paul George Ringo]

a := names[0:2]
b := names[1:3]
fmt.Println(a, b)   // [John Paul] [Paul George]

b[0] = "XXX"
fmt.Println(a, b)   // [John XXX] [XXX George]
fmt.Println(names)  // [John XXX George Ringo]
```

**关键点:**
- 切片就像数组的**引用**
- 更改切片的元素会修改其底层数组
- 共享底层数组的切片都会观测到这些修改
- 切片不存储数据,只是描述底层数组的一段

---

### 09_slice-literals

**切片字面量**

切片字面量类似于没有长度的数组字面量。

```go
// 数组字面量
[3]bool{true, true, false}

// 切片字面量 (会创建数组,然后构建引用它的切片)
[]bool{true, true, false}
```

**示例:**
```go
q := []int{2, 3, 5, 7, 11, 13}
r := []bool{true, false, true, true, false, true}

// 结构体切片
s := []struct {
    i int
    b bool
}{
    {2, true},
    {3, false},
    {5, true},
}
```

**关键点:**
- 切片字面量会先创建数组,再构建引用它的切片
- 语法类似数组字面量,但不指定长度
- 可以包含任意类型,包括结构体

---

### 10_slice-bounds

**切片的默认行为**

在进行切片时,可以利用默认行为来忽略上下界。

```go
s := []int{2, 3, 5, 7, 11, 13}

s = s[1:4]
fmt.Println(s)  // [3 5 7]

s = s[:2]
fmt.Println(s)  // [3 5]

s = s[1:]
fmt.Println(s)  // [5]
```

**默认值:**
- 下界默认值为 `0`
- 上界默认值为该切片的长度

**等价表达式:**
```go
var a [10]int

a[0:10]  // 完整形式
a[:10]   // 省略下界
a[0:]    // 省略上界
a[:]     // 省略上下界
```

---

### 11_slice-len-cap

**切片的长度与容量**

切片拥有**长度**和**容量**。

```go
s := []int{2, 3, 5, 7, 11, 13}
printSlice(s)  // len=6 cap=6 [2 3 5 7 11 13]

s = s[:0]
printSlice(s)  // len=0 cap=6 []

s = s[: 4]
printSlice(s)  // len=4 cap=6 [2 3 5 7]

s = s[2:]
printSlice(s)  // len=2 cap=4 [5 7]
```

**概念:**

| 属性 | 说明 | 获取方式 |
|------|------|---------|
| **长度** | 切片所包含的元素个数 | `len(s)` |
| **容量** | 从切片第一个元素到底层数组末尾的元素个数 | `cap(s)` |

**关键点:**
- 可以通过重新切片来扩展切片(前提是有足够容量)
- 截取切片会减少长度但保留容量
- 向后切片会同时减少长度和容量

---

### 12_nil-slices

**nil 切片**

切片的零值是 `nil`。

```go
var s []int
fmt.Println(s, len(s), cap(s))  // [] 0 0
if s == nil {
    fmt. Println("nil!")
}
```

**关键点:**
- nil 切片的长度和容量为 `0`
- nil 切片没有底层数组
- 可以用 `== nil` 判断切片是否为 nil

---

### 13_making-slices

**用 make 创建切片**

切片可以用内置函数 `make` 来创建,这也是创建动态数组的方式。

```go
a := make([]int, 5)       // len=5 cap=5 [0 0 0 0 0]
b := make([]int, 0, 5)    // len=0 cap=5 []
c := b[:2]                // len=2 cap=5 [0 0]
d := c[2:5]               // len=3 cap=3 [0 0 0]
```

**make 语法:**
```go
make([]T, len)       // 长度和容量都是 len
make([]T, len, cap)  // 长度为 len,容量为 cap
```

**关键点:**
- `make` 分配一个元素为零值的数组并返回引用它的切片
- 可以指定长度和容量
- 这是创建动态数组的标准方式

---

### 14_slices-of-slice

**切片的切片**

切片可以包含任何类型,包括其他切片。

```go
// 创建一个井字棋(经典游戏)
board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}

board[0][0] = "X"
board[2][2] = "O"
board[1][2] = "X"
board[1][0] = "O"
board[0][2] = "X"

for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
}
```

**输出:**
```
X _ X
O _ X
_ _ O
```

**关键点:**
- 切片可以嵌套形成多维数据结构
- 常用于矩阵、表格等二维数据
- 每个内层切片可以有不同的长度

---

### 15_append

**向切片追加元素**

为切片追加元素使用内置的 `append` 函数。

```go
func append(s []T, vs ...T) []T
```

**示例:**
```go
var s []int
printSlice(s)  // len=0 cap=0 []

s = append(s, 0)
printSlice(s)  // len=1 cap=1 [0]

s = append(s, 1)
printSlice(s)  // len=2 cap=2 [0 1]

s = append(s, 2, 3, 4)
printSlice(s)  // len=5 cap=6 [0 1 2 3 4]
```

**关键点:**
- `append` 的第一个参数是切片,其余是要追加的值
- 返回包含所有元素的新切片
- 当底层数组太小时,会分配更大的数组
- 可以一次追加多个元素
- 可在空切片 (nil) 上追加

**深入阅读:**
- [Go 切片: 用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals)

---

### 16_range

**range 遍历**

`for` 循环的 `range` 形式可遍历切片或映射。

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
}
```

**输出:**
```
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
... 
```

**关键点:**
- 每次迭代返回两个值:下标和元素的副本
- 第一个值:  当前元素的下标
- 第二个值: 该下标对应元素的**副本**

---

### 17_range-continued

**range 遍历(续)**

可以将下标或值赋予 `_` 来忽略它。

```go
pow := make([]int, 10)

// 只需要索引
for i := range pow {
    pow[i] = 1 << uint(i)  // == 2**i
}

// 只需要值
for _, value := range pow {
    fmt. Printf("%d\n", value)
}
```

**使用方式:**

| 语法 | 说明 |
|------|------|
| `for i, v := range slice` | 获取索引和值 |
| `for i, _ := range slice` | 只获取索引,忽略值 |
| `for _, v := range slice` | 只获取值,忽略索引 |
| `for i := range slice` | 只获取索引(简写) |

**关键点:**
- 使用 `_` 忽略不需要的值
- 若只需要索引,可以省略第二个变量

---

### 18_exercise-slices

**练习: 切片 - 生成图像**

实现 `Pic` 函数,返回一个二维切片表示图像。

```go
func Pic(dx, dy int) [][]uint8 {
    s := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        s[y] = make([]uint8, dx)
        for x := 0; x < dx; x++ {
            s[y][x] = uint8(x * int(math.Log(float64(y))))
        }
    }
    return s
}
```

**任务要求:**
- 返回长度为 `dy` 的切片
- 每个元素是长度为 `dx`,类型为 `uint8` 的切片
- 整数值会被解释为灰度值(蓝度值)

**有趣的函数:**
- `(x+y)/2`
- `x*y`
- `x^y`
- `x*log(y)`
- `x%(y+1)`

---

## 映射 (Map)

### 19_maps

**map 映射**

`map` 映射将键映射到值。

```go
type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{40.68433, -74.39967}
    fmt.Println(m["Bell Labs"])
}
```

**关键点:**
- 映射的零值为 `nil`
- `nil` 映射既没有键,也不能添加键
- `make` 函数返回初始化的映射

**语法:**
```go
var m map[KeyType]ValueType
```

---

### 20_map-literals

**映射字面量**

映射的字面量类似结构体,但必须有键名。

```go
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": Vertex{40.68433, -74.39967},
    "Google":     Vertex{37.42202, -122.08408},
}
```

**关键点:**
- 映射字面量必须指定键名
- 语法类似结构体字面量
- 可以在定义时初始化

---

### 21_map-literals-continued

**映射字面量(续)**

若顶层类型只是一个类型名,可以在字面量元素中省略它。

```go
var m = map[string]Vertex{
    "Bell Labs":  {40.68433, -74.39967},    // 省略 Vertex
    "Google":    {37.42202, -122.08408},   // 省略 Vertex
}
```

**关键点:**
- 当值类型明确时,可以省略类型名
- 使代码更简洁

---

### 22_mutating-maps

**修改映射**

映射的基本操作。

**操作:**

| 操作 | 语法 | 说明 |
|------|------|------|
| 插入/修改 | `m[key] = elem` | 设置元素 |
| 获取 | `elem = m[key]` | 读取元素 |
| 删除 | `delete(m, key)` | 删除元素 |
| 检测键存在 | `elem, ok = m[key]` | 双赋值检测 |

**示例:**
```go
m := make(map[string]int)

m["答案"] = 42
fmt. Println("值:", m["答案"])  // 值: 42

m["答案"] = 48
fmt.Println("值:", m["答案"])  // 值: 48

delete(m, "答案")
fmt.Println("值:", m["答案"])  // 值: 0

v, ok := m["答案"]
fmt.Println("值:", v, "是否存在?", ok)  // 值: 0 是否存在? false
```

**双赋值检测:**
- 若 `key` 在映射中,`ok` 为 `true`
- 若 `key` 不在映射中,`ok` 为 `false`,`elem` 为零值
- 可使用短变量声明:  `elem, ok := m[key]`

---

### 23_exercise-maps

**练习: 映射 - 单词计数**

实现 `WordCount`,返回字符串中每个单词的个数。

```go
func WordCount(s string) map[string]int {
    m := make(map[string]int)
    words := strings.Fields(s)  // 拆分为单词切片
    for _, word := range words {
        m[word]++
    }
    return m
}
```

**关键点:**
- 使用 `strings.Fields` 分割字符串
- 使用 `map[word]++` 统计单词出现次数
- 利用映射自动初始化为零值的特性

---

## 函数值与闭包

### 24_function-values

**函数值**

函数也是值,可以像其他值一样传递。

```go
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math. Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))    // 13
    
    fmt.Println(compute(hypot))   // 5
    fmt.Println(compute(math.Pow))  // 81
}
```

**关键点:**
- 函数可以作为值使用
- 函数值可以用作函数的参数或返回值
- 可以定义匿名函数并赋值给变量
- 函数类型:  `func(参数类型) 返回值类型`

---

### 25_function-closures

**函数闭包**

Go 函数可以是一个闭包。闭包是一个函数值,它引用了其函数体之外的变量。

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt. Println(pos(i), neg(-2*i))
    }
}
```

**输出:**
```
0 0
1 -2
3 -6
6 -12
10 -20
... 
```

**关键点:**
- 闭包是引用了外部变量的函数值
- 函数可以访问并赋予其引用的变量值
- 每个闭包都被绑定在其各自的变量上
- 闭包会"记住"创建时的环境

---

### 26_exercise-fibonacci-closure

**练习: 斐波纳契闭包**

实现一个返回斐波纳契数列的闭包函数。

```go
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt. Println(f())
    }
}
```

**输出:**
```
1
1
2
3
5
8
13
21
34
55
```

**关键点:**
- 使用闭包保存状态 (`a` 和 `b`)
- 每次调用返回下一个斐波纳契数
- 利用 Go 的多重赋值特性:  `a, b = b, a+b`
- [斐波纳契数列](https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97): 0, 1, 1, 2, 3, 5, 8, 13, ... 

---

## 核心概念对比

### 数组 vs 切片

| 特性 | 数组 | 切片 |
|------|------|------|
| 大小 | 固定 | 动态 |
| 类型 | `[n]T` | `[]T` |
| 长度是否为类型一部分 | 是 | 否 |
| 是否可改变大小 | 否 | 是 |
| 是否为引用类型 | 否(值类型) | 是(引用类型) |
| 常用程度 | 较少 | 非常常用 |

### 切片的内部结构

切片有三个组成部分:
1. **指针**: 指向底层数组
2. **长度**: 切片包含的元素个数
3. **容量**: 从切片第一个元素到底层数组末尾的元素个数

```
数组:  [2 3 5 7 11 13]
切片: [3 5 7]
      ↑     ↑
    指针  长度=3, 容量=5
```

### Map 特点

- 键必须是可比较的类型
- 零值为 `nil`
- 使用 `make` 创建
- 动态增长
- 无序
- 引用类型

---

## 运行示例

每个子目录都包含一个 `main.go` 文件,可以直接运行: 

```bash
# 运行某个示例
cd 01_pointers
go run main.go

# 或者直接指定路径
go run 01_pointers/main.go
```

**特殊示例** (需要外部包):
```bash
# 18_exercise-slices (图像生成)
go get golang.org/x/tour/pic
go run 18_exercise-slices/main.go

# 23_exercise-maps (单词计数测试)
go get golang.org/x/tour/wc
go run 23_exercise-maps/main.go
```

---

## 学习顺序建议

建议按照以下模块顺序学习:

1. **指针与结构体** (01-05)
   - 理解指针的基本概念
   - 掌握结构体的定义和使用
   - 学习结构体指针的隐式解引用

2. **数组与切片基础** (06-12)
   - 了解数组的局限性
   - 掌握切片的创建和使用
   - 理解切片作为引用的特性

3. **切片高级操作** (13-18)
   - 学习 `make` 和 `append`
   - 掌握 `range` 遍历
   - 完成切片练习

4. **映射 (Map)** (19-23)
   - 理解键值对数据结构
   - 掌握映射的增删改查
   - 完成映射练习

5. **函数高级特性** (24-26)
   - 理解函数作为值
   - 掌握闭包的概念和应用
   - 完成斐波纳契闭包练习

---

## 关键概念总结

### 指针
- `*T` 指向类型 `T` 的指针
- `&` 取地址,`*` 解引用
- Go 没有指针运算

### 结构体
- 使用 `type` 和 `struct` 定义
- 字段通过 `.` 访问
- 支持隐式指针解引用

### 数组
- 固定长度:  `[n]T`
- 长度是类型的一部分
- 很少直接使用

### 切片
- 动态大小:  `[]T`
- 引用类型,指向底层数组
- 有长度和容量
- 使用 `make` 创建,`append` 追加
- `range` 遍历

### 映射
- 键值对:  `map[K]V`
- 使用 `make` 创建
- 通过双赋值检测键是否存在
- 使用 `delete` 删除

### 函数值与闭包
- 函数是一等公民
- 可以作为参数和返回值
- 闭包捕获外部变量
- 保持状态

---

## 参考资源

- [Go 语言之旅 - 更多类型](https://tour.go-zh.org/moretypes/1) - 官方交互式教程
- [Go 切片: 用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals) - 深入理解切片
- [Effective Go - 数据](https://go.dev/doc/effective_go#data) - 官方最佳实践
- [Go 语言规范 - 类型](https://go.dev/ref/spec#Types) - 官方语言规范
- [斐波纳契数列](https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97) - 数学背景

---

**注意:** 本目录的示例代码展示了 Go 语言中最常用也最重要的数据结构和概念。特别是切片、映射和闭包,它们是 Go 编程的核心基础,务必深入理解和熟练掌握。