// 想要写出好的 Go 程序，单元测试是很重要的一部分。
// testing 包为提供了编写单元测试所需的工具，写好单元测试后，我们可以通过 go test 命令运行测试
// 为方便演示，例子的代码位于 main 包，实际上，单元测试的代码可以位于任何包下。
// 测试代码通常与需要被测试的代码位于同一个包下
package main

// 我们要测试下面这个简单的函数——返回最小值。
// 一般地，需要被测试的代码应该在类似于 intutils.go 的文件下， 其对应的测试文件应该被命名为 intutils_test.go
func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
