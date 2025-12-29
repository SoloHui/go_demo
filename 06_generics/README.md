# 06_generics - Go 泛型示例

本目录包含 Go 语言泛型(Generics)相关的示例代码,演示了类型参数和泛型类型的使用。

## 目录结构

- **01_index** - 类型参数示例
- **02_list** - 泛型类型示例

## 示例说明

### 01_index - 类型参数

演示如何使用类型参数编写可处理多种类型的 Go 函数。

**核心概念:**
- 类型参数出现在函数参数之前的方括号之间
- `comparable` 约束允许使用 `==` 和 `!=` 运算符
- 函数可以适用于任何支持比较的类型

**示例代码:**
```go
func Index[T comparable](s []T, x T) int {
    for i, v := range s {
        if v == x {
            return i
        }
    }
    return -1
}
```

**使用场景:**
- 在整数切片中查找元素
- 在字符串切片中查找元素
- 任何可比较类型的查找操作

### 02_list - 泛型类型

演示如何使用类型参数定义泛型数据结构,实现了一个可以保存任意类型值的单链表。

**核心概念:**
- 类型可以使用类型参数进行参数化
- 对实现通用数据结构非常有用
- 支持任意类型的存储和操作

**数据结构:**
```go
type List[T any] struct {
    next *List[T]
    val  T
}
```

**实现的方法:**
- `Append[T any](list *List[T], val T) *List[T]` - 向链表追加元素
- `PrintList[T any](list *List[T])` - 打印链表所有元素
- `Remove[T comparable](list *List[T], val T) *List[T]` - 删除指定值的节点
- `Set[T any](list *List[T], index int, val T)` - 设置指定索引的值

**使用示例:**
```go
// 整数链表
var intList *List[int]
intList = Append(intList, 10)
intList = Append(intList, 20)

// 字符串链表
var stringList *List[string]
stringList = Append(stringList, "foo")
stringList = Append(stringList, "bar")
```

## 运行示例

```bash
# 运行类型参数示例
cd 01_index
go run main. go

# 运行泛型类型示例
cd 02_list
go run main. go
```

## 参考资料

- [Go 语言之旅 - 泛型](https://tour.go-zh.org/generics/1)
- [Go 泛型官方文档](https://go.dev/doc/tutorial/generics)

## 学习要点

1. **类型参数** - 使函数可以处理多种类型
2. **类型约束** - 如 `comparable` 和 `any`
3. **泛型类型** - 创建可重用的数据结构
4. **类型推断** - Go 编译器可以自动推断类型参数

---

> 本示例来自 [Go 语言之旅](https://tour.go-zh.org/)