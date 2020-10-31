package main

import (
	"fmt"
	"sync"
)

var value int

//对value执行1000次++操作
//想对共享的变量进行修改时，需要先申请锁
//申请成功，可以修改，如果锁已经被别人使用，那么阻塞等待
//RwMutx 读写锁 允许多个读goroutine共享读锁
// var rwlock sync.RWMutex
// rwlock.
// Cond
var cond sync.Cond

func add() {
	mutex.Lock()
	value++ //并发会导致此行代码发生问题
	mutex.Unlock()
	wg.Done()
}

var wg sync.WaitGroup
var mutex sync.Mutex //锁
//mutex.Lock() 申请锁，如果成功，则继续执行
//mutex.Unlock() 释放锁，一定现拥有锁，才可以释放，释放后可以帮助其他阻塞的goroutine解除阻塞

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(value)
}
