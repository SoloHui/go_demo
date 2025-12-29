# Go 语言方法与接口 (05_methods)

本目录包含 Go 语言中方法 (Methods) 和接口 (Interfaces) 的学习示例代码。主要参考自 [Go 语言之旅](https://tour.go-zh.org/)。

## 目录结构

### 方法 (Methods)
- [01_methods](#01_methods) - 方法定义
- [02_methods-funcs](#02_methods-funcs) - 方法即函数
- [03_methods-continued](#03_methods-continued) - 非结构体类型的方法
- [04_methods-pointers](#04_methods-pointers) - 指针接收者
- [05_methods-pointers-explained](#05_methods-pointers-explained) - 指针接收者详解
- [06_indirection](#06_indirection) - 方法与指针重定向
- [07_indirection-values](#07_indirection-values) - 方法与值重定向
- [08_methods-with-pointer-receivers](#08_methods-with-pointer-receivers) - 选择值或指针作为接收者

### 接口 (Interfaces)
- [09_interfaces](#09_interfaces) - 接口定义
- [10_interfaces-are-satisfied-implicitly](#10_interfaces-are-satisfied-implicitly) - 隐式实现接口
- [11_interface-values](#11_interface-values) - 接口值
- [12_interface-values-with-nil](#12_interface-values-with-nil) - 底层值为 nil 的接口值
- [13_nil-interface-values](#13_nil-interface-values) - nil 接口值
- [14_empty-interface](#14_empty-interface) - 空接口
- [15_type-assertions](#15_type-assertions) - 类型断言
- [16_type-switches](#16_type-switches) - 类型选择

### 常用标准库接口
- [17_stringer](#17_stringer) - Stringer 接口
- [18_exercise-stringer](#18_exercise-stringer) - 练习: Stringer
- [19_errors](#19_errors) - Errors 接口
- [20_exercise-errors](#20_exercise-errors) - 练习: Errors
- [21_reader](#21_reader) - Reader 接口
- [22_exercise-reader](#22_exercise-reader) - 练习: Reader
- [23_exercise-rot-reader](#23_exercise-rot-reader) - 练习: rot13Reader
- [24_images](#24_images) - Image 接口
- [25_exercise-images](#25_exercise-images) - 练习: Images

---

## 方法 (Methods)

### 01_methods

**方法定义**

Go 没有类。不过你可以为类型定义方法。
方法就是一类带特殊的 **接收者** 参数的函数。

```go
type Vertex struct {
    X, Y float64
}

// Abs 方法拥有一个名字为 v，类型为 Vertex 的接收者
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

**关键点:**
- 方法接收者在 `func` 关键字和方法名之间
- 可以为任意类型定义方法(不仅仅是结构体)

---

### 04_methods-pointers

**指针接收者**

你可以为指针类型的接收者声明方法。
这意味着对于某类型 `T`，接收者的类型可以用 `*T` 的文法。

```go
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
```

**使用指针接收者的原因:**
1. **修改接收者**: 方法能够修改其接收者指向的值。
2. **避免复制**: 避免在每次调用方法时复制该值(对于大型结构体更高效)。

**关键点:**
- 指针接收者的方法可以修改接收者指向的值
- 若使用值接收者，方法会对原始值的副本进行操作

---

### 08_methods-with-pointer-receivers

**选择值或指针作为接收者**

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。

```go
func (v *Vertex) Scale(f float64) {
    // ...
}

func (v *Vertex) Abs() float64 {
    // ...
}
```

**关键点:**
- 如果一个方法需要指针接收者，通常所有方法都应该使用指针接收者
- 这样可以保持方法集的一致性

---

## 接口 (Interfaces)

### 09_interfaces

**接口定义**

接口类型是由一组方法签名定义的集合。
接口类型的变量可以保存任何实现了这些方法的值。

```go
type Abser interface {
    Abs() float64
}
```

**关键点:**
- 接口定义了一组行为
- 类型通过实现接口中的方法来实现接口

---

### 10_interfaces-are-satisfied-implicitly

**隐式实现**

类型通过实现一个接口的所有方法来实现该接口。
既然无需专门显式声明，也就没有 "implements" 关键字。

```go
type I interface {
    M()
}

type T struct {
    S string
}

// T 实现了接口 I，但无需显式声明
func (t T) M() {
    fmt.Println(t.S)
}
```

**关键点:**
- 隐式实现解耦了实现接口的包和定义接口的包
- 鼓励定义精确的、小型的接口

---

### 14_empty-interface

**空接口**

指定了零个方法的接口值被称为 *空接口*：

```go
interface{}
```

空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

```go
var i interface{}
i = 42
i = "hello"
```

**关键点:**
- 空接口用于处理未知类型的值
- 类似于 Java 中的 Object 或 C 中的 void*

---

### 15_type-assertions

**类型断言**

类型断言提供了访问接口值底层具体值的方式。

```go
t := i.(T)      // 断言 i 保存了 T 类型，若不是则 panic
t, ok := i.(T)  // 安全断言，ok 为 false 时不 panic
```

**示例:**
```go
var i interface{} = "hello"

s := i.(string)
fmt.Println(s)

s, ok := i.(string)
fmt.Println(s, ok)

f, ok := i.(float64)
fmt.Println(f, ok)
```

**关键点:**
- 使用 `.(T)` 语法
- 推荐使用双返回值形式 `t, ok := i.(T)` 以避免 panic

---

### 16_type-switches

**类型选择**

类型选择是一种按顺序从几个类型断言中选择分支的结构。

```go
switch v := i.(type) {
case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
default:
    fmt.Printf("I don't know about type %T!\n", v)
}
```

**关键点:**
- `switch v := i.(type)` 语法
- `case` 中使用类型而非值

---

## 常用标准库接口

### 17_stringer

**Stringer**

`fmt` 包中定义的 `Stringer` 是最普遍的接口之一。

```go
type Stringer interface {
    String() string
}
```

**示例:**
```go
func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```

**关键点:**
- 实现 `String()` 方法来自定义打印格式
- 类似于 Java 的 `toString()`

---

### 19_errors

**Errors**

Go 程序使用 `error` 值来表示错误状态。

```go
type error interface {
    Error() string
}
```

**用法:**
```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
```

**关键点:**
- `error` 是一个内置接口
- 通常作为函数的最后一个返回值
- 检查 `err != nil` 来处理错误

---

### 21_reader

**Readers**

`io` 包指定了 `io.Reader` 接口，它表示数据流的读取端。

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 `io.EOF` 错误。

---

### 24_images

**Images**

`image` 包定义了 `Image` 接口：

```go
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

---

## 运行示例

每个子目录都包含一个 `main.go` 文件，可以直接运行：

```bash
# 运行某个示例
cd 01_methods
go run main.go

# 或者直接指定路径
go run 01_methods/main.go
```

**特殊示例** (需要外部包):
```bash
# 22_exercise-reader
go get golang.org/x/tour/reader

# 24_images
go get golang.org/x/tour/pic
```

---

## 学习顺序建议

1. **方法基础** (01-08)
   - 理解方法与函数的区别
   - 重点掌握指针接收者与值接收者的区别

2. **接口基础** (09-16)
   - 理解接口的隐式实现机制
   - 掌握空接口、类型断言和类型选择

3. **标准库接口** (17-25)
   - 熟练使用 `Stringer` 和 `error`
   - 了解 `io.Reader` 的工作模式

---

## 关键概念总结

### 方法
- 方法是带接收者的函数
- 接收者可以是值或指针
- 指针接收者可修改原值，且避免拷贝

### 接口
- 接口定义方法集合
- 隐式实现（Duck Typing）
- `interface{}` 可接收任意类型
- 类型断言 `x.(T)` 用于提取具体类型

### 常用接口
- `Stringer`: 自定义打印格式
- `error`: 错误处理
- `Reader`: 数据流读取

---

## 参考资源

- [Go 语言之旅 - 方法](https://tour.go-zh.org/methods/1)
- [Effective Go - 方法](https://go.dev/doc/effective_go#methods)
- [Effective Go - 接口](https://go.dev/doc/effective_go#interfaces)
