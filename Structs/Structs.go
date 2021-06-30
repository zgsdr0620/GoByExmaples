// Go 的结构体(struct) 是带类型的字段(fields)集合。 这在组织数据时非常有用
package main

import "fmt"

// 这里的 person 结构体包含了 name 和 age 两个字段
type person struct {
	name string
	age  int
}

// newPerson根据给定的name构建一个新的person结构体
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// 您可以安全地返回指向局部变量的指针，因为局部变量将在函数作用域内继续存在
	return &p
}

func main() {
	// 使用这个语法创建新的结构体元素
	fmt.Println(person{"Bob", 20})

	// 你可以在初始化一个结构体元素时指定字段名字
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的字段将被初始化为零值
	fmt.Println(person{name: "Fred"})

	// & 前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	// 在构造函数中封装新的结构体创建是惯用的
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	// 使用.来访问结构体字段
	fmt.Println(s.name)

	sp := &s
	// 也可以对结构体指针使用.
	// 指针会被自动解引用
	fmt.Println(sp.age)

	// 结构体是可变(mutable)的
	sp.age = 51
	fmt.Println(sp.age)
}
