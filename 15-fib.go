//计算斐波那契数列

package main

import (
	"fmt"
	"time"
)

/*
 	x = 0 , f(x) = 0
	x = 1 , f(x) = 1
	x = 2, f(x) = f(0) + f(1)
	f(N) = f(N-1) + f(N-2)
*/

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//提升用户体验
func deday() {
	//_,v 这种方式代表忽略第一个返回值,_代表占位
	val := "\\|/-"
	for {

		// for _, v := range `\|/-` {
		// 	fmt.Printf("\r%c", v)
		// 	time.Sleep(time.Millisecond * 100) //每隔0.1s打印一下
		// }

		for _, v := range []byte(val) {
			fmt.Printf("\r%c", v)
			time.Sleep(time.Millisecond * 100) //每隔0.1s打印一下
		}
	}
}

func main() {
	go deday()
	fmt.Println(fib(45))
}
