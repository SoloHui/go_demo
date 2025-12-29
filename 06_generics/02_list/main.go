/*
---
title: Go 语言之旅
url: https://tour.go-zh.org/generics/2
publishedTime: undefined
---

## 泛型类型

除了泛型函数之外，Go 还支持泛型类型。 类型可以使用类型参数进行参数化，这对于实现通用数据结构非常有用。

此示例展示了能够保存任意类型值的单链表的简单类型声明。

作为练习，请为此链表的实现添加一些功能。
*/
package main

import "fmt"

// List 表示一个可以保存任何类型的值的单链表。
type List[T any] struct {
	next *List[T]
	val  T
}

// 练习：为 List 实现一些方法。
func Append[T any](list *List[T], val T) *List[T] {
	newList := &List[T]{val: val}
	if list == nil {
		return newList
	}
	curr := list
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newList
	return list
}

func PrintList[T any](list *List[T]) {
	curr := list
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
}

func Remove[T comparable](list *List[T], val T) *List[T] {
	if list == nil {
		return nil
	}
	if list.val == val {
		return list.next
	}
	curr := list
	for curr.next != nil {
		if curr.next.val == val {
			curr.next = curr.next.next
			return list
		}
		curr = curr.next
	}
	return list
}

func Set[T any](list *List[T], index int, val T) {
	curr := list
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.next
	}
	if curr != nil {
		curr.val = val
	}
}

func main() {
	var intList *List[int]
	intList = Append(intList, 10)
	intList = Append(intList, 20)
	intList = Append(intList, 30)
	fmt.Println("Integer List:")
	PrintList(intList)

	intList = Remove(intList, 20)
	fmt.Println("After removing 20:")
	PrintList(intList)

	Set(intList, 1, 99)
	fmt.Println("After setting index 1 to 99:")
	PrintList(intList)

	var stringList *List[string]
	stringList = Append(stringList, "foo")
	stringList = Append(stringList, "bar")
	stringList = Append(stringList, "baz")
	fmt.Println("String List:")
	PrintList(stringList)

	stringList = Remove(stringList, "bar")
	fmt.Println("After removing 'bar':")
	PrintList(stringList)

	Set(stringList, 0, "qux")
	fmt.Println("After setting index 0 to 'qux':")
	PrintList(stringList)
}
