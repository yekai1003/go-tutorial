package main

import (
	"fmt"
	"time"
)

var ch chan string //chan string 代表channel类型

func write() {
	fmt.Println("I am wirte goroutine")
	time.Sleep(time.Second * 3) //睡眠3s的目的是为了感受main-读的阻塞
	ch <- "heihei---------"     //写入channel
	fmt.Println("begin write twice")
	ch <- "haha---------" //写入channel
	fmt.Println("after write")
}

func main() {
	//创建channel 使用make
	//ch := make(chan string)
	ch = make(chan string)
	//ch = make(chan string, 10) //带缓冲区的channel
	go write()
	fmt.Println("I am main goroutine")
	msg := <-ch //读channel
	fmt.Println("after read,msg = ", msg)
	time.Sleep(time.Second * 3)
	msg = <-ch //读channel
	fmt.Println("after read,msg = ", msg)

}
