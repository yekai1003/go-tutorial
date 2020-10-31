/*

	容器类型，引用类型
*/
package main

import (
	"fmt"
)

func main() {
	//定义数组,数组也是变量
	var a1 [5]int
	var a2 []int = []int{1, 2, 3, 4, 5} //元素个数不能超过数组上限
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a2[2])
	//循环关键字只有一个for
	for i := 0; i < 5; i++ {
		fmt.Println(a2[i])
	}
	//切片 可以从数组或已经存在的切片中抽取内容,动态数组
	//切片语法  array|slice[begin:end] -- 区间[begin, end), [begin, end-1]
	//begin 和end可以省略
	//s1 是一个变量
	s1 := a2[2:4]
	fmt.Println(s1)
	fmt.Println(a2[:4])
	fmt.Println(a2[2:])
	s1[0] = 100
	fmt.Println("----------------")
	fmt.Println(s1)
	fmt.Println(a2) //切片修改，数组受影响,切片仍在数组原地址
	//向切片追加元素 append
	s1 = append(s1, 1000)
	s1 = append(s1, 10000)
	s1 = append(s1, 100000)
	// s2 := append(s1, 1000)
	fmt.Println(s1)
	fmt.Println(a2)
	s1[0] = 30000
	fmt.Println(s1)
	fmt.Println(a2)
	// len cap
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(a2), cap(a2))
	s1 = append(s1, 100000)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 100000)
	fmt.Println(len(s1), cap(s1))
	//copy
	var s2 []int
	s2 = append(s2, 100, 200, 300)
	fmt.Println(len(s2), cap(s2))
	//切片在使用时需要构建 make ,比如拷贝，或需要系统调用写入时
	//s2 = make([]int, 3, 8) // make(T, len[, cap])
	fmt.Println(s2)
	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))
}
