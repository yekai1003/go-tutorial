//TCP服务器
//TCP想要建立连接，需要三次握手
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//1. 建立一个侦听器 Listen,需要tcp和地址
	//地址： ip + 端口 ，共同描述为一个进程（唯一一个主机上的）

	lnr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic("Failed to Listen", err)
	}
	defer lnr.Close()
	//2. 借助侦听器获取新的连接
	//从已连接队列中获取一条连接信息
	conn, err := lnr.Accept()
	//3. 与连接进行通信
	buf := make([]byte, 256)
	fmt.Println("new conn:", conn.RemoteAddr())
	conn.Read(buf)
	conn.Write(buf) //写给网络
	os.Stdout.Write(buf)
}
