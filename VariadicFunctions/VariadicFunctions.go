package main

import "fmt"

// 可变参数函数。 在调用时可以传递任意数量的参数。 例如，fmt.Println 就是一个常见的变参函数

// 这个函数接受任意数量的 int 作为参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}
func main() {

	sum(1, 2)
	sum(1, 2, 3)

	// 如果你有一个含有多个值的 slice，想把它们作为参数使用， 你需要这样调用 func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
