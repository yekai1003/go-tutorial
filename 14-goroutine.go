package main

import (
	"fmt"
	"time"
)

/*
	并发：同一时间间隔内，多个进程在同时运行
		多进程或多线程
	并行：同一时刻，多个进程在同时运行
	Go语言的并发最小粒度是goroutine，例程，协程
	语法层面支持，自己开发了调度器
	线程 - Context（处理器，上下文） - goroutine列表
	启动goroutine: go funcname(...)
*/

func demo() {
	fmt.Println("I am a goroutine")
}

func main() {
	go demo() //启动一个
	fmt.Println("I am main goroutine")
	//main函数退出，整个进程就退出
	time.Sleep(time.Second * 1) //睡眠1s，Sleep不必关心类型是什么，直接用宏
}
