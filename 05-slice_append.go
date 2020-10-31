package main

import (
	"fmt"
)

func main() {
	var s1 []int
	fmt.Println(len(s1), cap(s1))
	for i := 0; i < 20; i++ {
		s1 = append(s1, i)
		fmt.Println(len(s1), cap(s1))
	}
}
