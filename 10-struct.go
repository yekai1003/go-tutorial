package main

import (
	"fmt"
)

/*
	语法：type StructName struct {
		filedName1 T1
		filedName2 T2
	}
*/
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p1 Person = Person{"yekai", 30, "man"}
	p2 := Person{
		name: "xiaohong",
		age:  30, // ,必须加
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("%+v\n", p2)
}
