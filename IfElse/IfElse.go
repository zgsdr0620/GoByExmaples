package main

import "fmt"

func main() {

	if 7%2 == 0 { // if + else
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // if
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // 带初始化变量声明的if + else if + else
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
