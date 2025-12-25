package main

// import 的作用是引入其他包，以便在当前文件中使用它们提供的功能。
// 这里我们引入了 fmt 包，它提供了格式化输入输出的功能。
import "fmt"

func main() {
	// fmt.Println 是 fmt 包中的一个函数，用于在控制台打印输出内容。
	// 这里我们打印了一条简单的欢迎信息。
	fmt.Println("Hello, welcome to Go programming!")

	// 我们还可以使用 fmt 包的其他函数，例如 fmt.Print。
	// 区别: fmt.Print 不会在输出末尾添加换行符。
	fmt.Print("This is a simple demonstration of Go's import and fmt package.\n")

	// fmt.Printf 用于格式化输出，可以指定输出的格式。
	// 这里我们演示了如何使用格式化字符串输出一个整数。
	fmt.Printf("This is a formatted number: %d\n", 42)

	/*
		多行注释
	*/
}
