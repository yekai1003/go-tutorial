package main

import (
	"fmt"
)

/*

	函数闭包：在一个函数内部，包含子函数，字函数可以访问该函数的局部变量
*/

func getNextNumber() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f := getNextNumber() //f 是一个函数指针， func() int
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	ff := getNextNumber()
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(f())
}
