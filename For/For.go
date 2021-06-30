package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { // 没有局部变量初始化 + 有循环条件
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ { // 有局部变量初始化 + 有循环条件
		fmt.Println(j)
	}

	for { // 没有局部变量初始化 + 没有循环条件，死循环
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { // 有局部变量初始化 + 有循环条件
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
