package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 6)
	go func() {
		//模拟写操作
		for i := 0; i < 10; i++ {
			fmt.Println("write ", i)
			ch <- i
			time.Sleep(time.Second * 1)

		}
	}()

	time.Sleep(time.Second * 10)
	for i := 0; i < 10; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}
