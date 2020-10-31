//火箭发射

package main

import (
	"fmt"
	"os"
	"time"
)

//模拟发射
func lanuch() {
	fmt.Println("火箭发射")
}

var ch chan int

//倒数计时，5，4，3，2，1 后 发射
//可以突然取消发射 通过按下回车取消
//解决思路 启动一个goroutine监控键盘输入，如果突然按下，通过channel通知给main

func cancel() {
	//os.Stdin 标准输入对象 Read读标准输入
	buf := make([]byte, 10)
	os.Stdin.Read(buf) //阻塞读
	//取消发射- 通知
	ch <- 0 //通知main
}

func main() {
	ch = make(chan int)
	tk := time.NewTicker(time.Second * 1) //定义每隔1s的定时器

	// for i := 5; i > 0; i-- {
	// 	<-tk.C
	// 	fmt.Println(i)
	// }
	//1个goroutine同时要读取2个channel容易死锁
	//select 可以监控多个channel，阻塞提前到select,任何一个chanel来消息都可以解除阻塞
	flag := true
	go cancel()
	num := 5
	for {
		select {
		case <-ch:
			flag = false
		case <-tk.C:
			fmt.Println(num)
			num--

		}
		if num == 0 || !flag {
			break
		}
	}
	if flag {
		lanuch()
	}

}
