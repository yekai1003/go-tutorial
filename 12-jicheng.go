/*
	继承
*/
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  string
}

//方法封装注意事项：如果要修改成员变量，那么不能值传递
func (p *Person) setAge(a int) {
	p.age = a
	fmt.Println("------", p)
}

type SuperMan struct {
	Person //继承，内嵌
	fight  uint
}

type SuperMan2 struct {
	Person //继承，内嵌
	fight  uint
	name   string
}

func main() {
	s1 := SuperMan{Person{"yekai", 30, "man"}, 1000000}
	fmt.Println(s1)
	p1 := Person{"yekai", 30, "man"}
	p1.setAge(100)
	fmt.Println(p1)
	//s1.setAge(300)
	s1.setAge(300)
	fmt.Println(s1)
	s2 := SuperMan2{Person{"yekai", 30, "man"}, 1000000, "fuxue"}
	fmt.Println(s2)
	fmt.Println(s2.Person.name)
}
