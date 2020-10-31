package main

import (
	"fmt"
)

//实现3个goroutine之间消息传递，第一个传递1，2，3，4，。。9给第2个，
//第2个将数字平方给第3个，第3个打印
// 用一个channel行不行？ 2个
// 每个goroutine操作channel都应该是单方向的
// 为了安全，使用单方向channel，仅限于函数参数部分

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//channel关闭 close(ch)

	//3个goroutine
	go func(in chan<- int) {
		//传递1，2给第2个
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in) //一定要写端关
	}(ch1)
	go func(out <-chan int, in chan<- int) {
		//接收1，2，并且平方后给第3个
		// for i := 0; i < 9; i++ {
		// 	val := <-ch1
		// 	ch2 <- val * val
		// }

		for {
			val, ok := <-out
			//ok 指示channel是否可读
			if !ok {
				break
			}
			in <- val * val
		}
		close(in)
	}(ch1, ch2)
	//main 作为第3个
	// for i := 0; i < 9; i++ {
	// 	msg := <-ch2
	// 	fmt.Println(msg)
	// }
	for {
		msg, ok := <-ch2
		if !ok {
			break
		}
		fmt.Println(msg)
	}
}
