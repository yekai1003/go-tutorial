package main

import (
	"fmt"
)

/*
	多态，首先要继承
	多态的前提：
	 	1. 向上赋值
		2. 调用方法要使用派生类的方法
*/
/*
	定义接口，支持多态
	type interfaceName interface {
	 	method1()
		method2()
	}
	接口定义方法后无需实现
*/
type Ainimal interface {
	sleep()
	eating()
}

type Dog struct {
	color string
}

//如果结构体实现了某接口内的全部方法，那么就代表结构体支持了该接口，反过来，
//接口类型对象可以指向结构体对象
func (d Dog) sleep() {
	fmt.Printf("%s dog is sleep\n", d.color)
}

// func (d Dog) eating() {
// 	fmt.Printf("%s dog is eating\n", d.color)
// }

type Cat struct {
	color string
}

func (c Cat) sleep() {
	fmt.Printf("%s cat is sleep\n", c.color)
}
func (c Cat) eating() {
	fmt.Printf("%s cat is eating\n", c.color)
}

func callDemo(a Ainimal) {
	a.eating()
	a.sleep()
}

func main() {
	d1 := Dog{"black"}
	//d1.eating()
	d1.sleep()
	c1 := Cat{"white"}
	c1.eating()
	var a1 Ainimal //定义接口
	a1 = d1
	a1.eating()
	a1 = c1
	a1.sleep()
	callDemo(c1)
	callDemo(d1)
}
