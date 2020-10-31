/*
	常量说明 const
*/
package main

import (
	"fmt"
)

const (
	pi  = 3.14
	str = "hello"
)

//定义枚举
// const (
// 	apple  = 0
// 	orange = 1
// 	banana = 2
// 	peach  = 3
// )

const (
	apple, c1 = iota, iota + 1 // iota初始化为0，按行自增
	orange, c2
	banana, c3 = iota * 2, iota + 3
	peach, c4
)

func main() {
	fmt.Println(pi, str)
	fmt.Println(apple, orange, banana, peach)
	fmt.Println(c1, c2, c3, c4)
}
