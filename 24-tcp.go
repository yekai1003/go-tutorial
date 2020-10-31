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
	for {
		//从已连接队列中获取一条连接信息
		conn, err := lnr.Accept()
		if err != nil {
			fmt.Println("failed to Accept")
			continue
		}
		go client_conn(conn)
	}

}

func client_conn(conn net.Conn) {
	//3. 与连接进行通信
	defer conn.Close()
	fmt.Println("new conn:", conn.RemoteAddr())

	buf := make([]byte, 256)
	for {

		n, _ := conn.Read(buf) //Read将系统缓冲区的内容读取并写入到buf中

		//1234567890
		//heihei
		conn.Write(buf[:n]) //写给网络 ,实际是将buf中的内容写到系统的写缓冲区中
		os.Stdout.Write(buf[:n])
	}

}
