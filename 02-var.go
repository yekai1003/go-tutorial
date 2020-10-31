/*
	变量知识点说明
*/
package main

import (
	"fmt"
)

func main() {
	//变量定义的方式: var varname  varType
	var x1 int // int , uint string, float32 ,float64, byte, bool
	fmt.Println(x1)
	//定义的时候同时初始化: var varname  varType = value
	var x2 float64 = 3.14
	fmt.Println(x2)
	// 第三种方式 := 自动推导变量类型，定义时，变量必须是当前代码段第一次使用,而且只能在函数内
	str3 := "yekai"
	fmt.Println(str3)
	str3 = "fuhongxue"
	//同时定义多个变量
	x, y := 100, "hello"
	fmt.Println(x, y)
}
