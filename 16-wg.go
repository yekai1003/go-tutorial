package main

import (
	"fmt"
	"sync"
)

/*
	并发：同一时间间隔内，多个进程在同时运行
		多进程或多线程
	并行：同一时刻，多个进程在同时运行
	Go语言的并发最小粒度是goroutine，例程，协程
	语法层面支持，自己开发了调度器
	线程 - Context（处理器，上下文） - goroutine列表
	启动goroutine: go funcname(...)
	sync包下的同步工具
	WaitGroup
	Mutex
	Rwmutx
	Cond
	Once
*/

func demo() {
	fmt.Println("I am a goroutine")
	wg.Done() //减少一个计数
}

//定义一个WaitGroup变量
var wg sync.WaitGroup

// wg.Add(delta int) 添加要监控的数量 ++
// wg.Done() 减少一个监控计数,执行完的gouroutine负责调用
//wg.Wait() 阻塞等待监控计数变为0，之后解除阻塞

func main() {
	wg.Add(1) // 增加一个计数
	go demo() //启动一个
	fmt.Println("I am main goroutine")
	wg.Wait() //阻塞等待监控的计数变为0

	//main函数退出，整个进程就退出
	//很希望demo（）结束后再退出
}
