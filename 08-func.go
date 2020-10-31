/*
	函数
*/
package main

import (
	"fmt"
)

/*
	函数声明：
	func func_name(paramlist) return list {
		//函数体
	}
	参数，注意类型后置
	return list 代表返回列表，可以返回多个值或0个值
*/

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

//定义一个统一的函数，将方法和函数一起传递
//func(a, b int) int  -- 整体代表函数类型
func callfunc(a, b int, f func(x, y int) int) int {
	return f(a, b)
}

//交换2个整形数
//值传递，交换无效
func swap(a, b int) {
	temp := a
	a = b
	b = temp
}

//使用指针
func swap2(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func swap3(a, b int) (int, int) {
	return b, a
}

func main() {
	x, y := 10, 20
	fmt.Println(add(x, y))
	f1 := add //函数赋值
	fmt.Println(f1(10, 30))
	fmt.Println(callfunc(20, 50, add))
	fmt.Println(callfunc(20, 50, sub))
	//匿名函数
	func(a, b int) {
		fmt.Printf("%d + %d = %d\n", a, b, a+b)
	}(10, 20)
	fmt.Println("-------------")
	swap(x, y)
	fmt.Println(x, y)
	swap2(&x, &y) //按照地址传入
	fmt.Println(x, y)
	x, y = swap3(x, y)
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

}
