package main

import "fmt"

func main() {
	// 1. 无初始化, 指定变量类型
	var name string
	var age int
	fmt.Printf("%q %d\n", name, age)
	// 输出结果
	// "" 0

	// 2. var 类型推断
	var name2 = "gopher2"
	fmt.Printf("%s\n", name2)
	// 输出结果
	// gopher2

	// 3. 短变量声明 :=
	name3 := "gopher3"
	fmt.Printf("%s\n", name3)
	// 输出结果
	// gopher3
}