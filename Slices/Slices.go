package main

import "fmt"

func main() {

	// 与数组不同，slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）
	// 要创建一个长度不为 0 的空 slice，需要使用内建函数 make
	// 这里我们创建了一个长度为 3 的 string 类型的 slice（初始值为零值）
	s := make([]string, 3)
	fmt.Println(s)

	// 我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len 返回 slice 的长度
	fmt.Println("len: ", len(s))


	// 除了基本操作外，slice 支持比数组更丰富的操作
	// 比如 slice 支持内建函数 append， 该函数会返回一个包含了一个或者多个新值的 slice
	// 注意由于 append 可能返回一个新的 slice，我们需要接收其返回值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slice 还可以 copy
	// 这里我们创建一个空的和 s 有相同长度的 slice——c， 然后将 s 复制给 c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice 支持通过 slice[low:high] 语法进行“切片”操作
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中声明并初始化一个 slice 变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// 内部的 slice 长度可以不一致，这一点和多维数组不同
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	// 注意，slice 和数组是不同的类型，但它们通过 fmt.Println 打印的输出结果是类似的
	fmt.Println("2d: ", twoD)
}
