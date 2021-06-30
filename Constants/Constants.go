package main

import (
	"fmt"
	"math"
)

const s string = "constant" // 指定类型 + 全局常量
func main() {
	fmt.Println(s)

	const n = 5000000000 // 未指定类型 + 局部常量

	const d = 3e20 / n // 未指定类型 + 局部常量
	fmt.Println(d)

	fmt.Println(int64(d)) // 强制类型转换

	fmt.Println(math.Sin(n))

}
