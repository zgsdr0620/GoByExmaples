package main

import "fmt"

func main() {

	var a = "initial" // 不指定类型 + 初始化
	fmt.Println(a)

	var b, c = 1, 2 // 多变量声明 + 指定类型 + 初始化
	fmt.Println(b, c)

	var d = true // 不指定类型 + 初始化
	fmt.Println(d)

	var e int // 指定类型 + 不初始化
	fmt.Println(e)

	f := "apple" // 短变量声明
	fmt.Println(f)

}
