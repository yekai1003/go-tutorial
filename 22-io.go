/*
	文件描述符： Linux系统的，进程内部会为打开的文件创建一个文件描述相关的表，
	文件描述符指向这个表里的一条信息
	掌握了文件描述符就可以对文件进行读写操作

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	//1.  打开文件
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 0666)
	// ~umask & 0666 == >
	/*
		root:godev yk$ umask
		0022
		0666 - 0022 = 0 644
		110 110 110 ==> 110 100 100 644
		000 010 010
		110 101 110 => 110 101 100
		000 010 010
	*/
	if err != nil {
		fmt.Println("Failed to OpenFile", err)
		return
	}
	defer f.Close() //defer可以让之后的函数调用延迟执行,等待代码段结束
	f.Write([]byte("heihei"))

}
