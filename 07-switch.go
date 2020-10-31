package main

import (
	"fmt"
)

func main() {

	//expr := "yekai"
	var expr string
	fmt.Scanf("%s", &expr) //从标准输入读入字符串，填写到expr中

	switch expr {
	case "yekai":
		fmt.Println("hello, yekai")
	case "xiaohong":
		fmt.Println("byebye, xiaohong")
	case "xiaowang":
		fmt.Println("???, xiaowang")
	default:
		fmt.Println("sorry, --------")
	}

	var number int
	fmt.Println("Please input a number:")
	fmt.Scanf("%d", &number)
	switch number {
	case 100:
		fmt.Println("very good ")
		fallthrough //可以执行之后的分支
	case 85:
		fmt.Println("very well")
		fallthrough
	case 60:
		fmt.Println("jige")
		fallthrough
	default:
		fmt.Println("sorry, --------")
	}
}
