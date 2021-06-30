package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	// 一个基本的 switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在同一个 case 语句中，你可以使用逗号来分隔多个表达式
	// 在这个例子中，我们还使用了可选的 default 分支
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	// 不带表达式的 switch 是实现 if/else 逻辑的另一种方式
	switch {
	// 这里还展示了 case 表达式也可以不使用常量
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			// 注意这里是Printf而不是Println
			// 后者不能格式化输出
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
