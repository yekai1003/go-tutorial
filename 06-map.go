/*

	键值对 map
*/
package main

import (
	"fmt"
	"reflect"
)

//nil 代表go语言的NULL
func main() {
	//1. 创建map
	// var score map[string]uint // map的key是string类型，value是uint类型
	// //map必须创建一下
	// score = make(map[string]uint)
	score := make(map[string]uint)
	//2. 添加内容
	score["yekai"] = 100
	score["fuhongxue"] = 150
	score["lixunhuan"] = 150
	score["afei"] = 150
	//3. 访问map元素
	fmt.Println(score)
	fmt.Println(score["yekai"])
	val, ok := score["yekai"]
	if ok {
		fmt.Println(val, ok)
	} else {
		fmt.Println("person not exists")
	}
	//对map遍历？ 地址不连续  range,会查找它之后的容器的所有元素，并且获取k-v值
	for k, v := range score {
		fmt.Println(k, "->", v)
	}
	s1 := []int{1, 2, 3, 4, 5, 6}
	for k, v := range s1 {
		fmt.Println(k, "->", v)
	}
	//接口类型 interface{} 可以理解为范型
	mapdemo := make(map[string]interface{})
	mapdemo["yekai"] = "yekai2"
	mapdemo["fuhongxue"] = 100
	mapdemo["zhangsan"] = 3.14
	fmt.Println(mapdemo)
	var str string
	//str = mapdemo["yekai"].(reflect.TypeOf(mapdemo["yekai"]).String()) //类型断言
	fmt.Println(reflect.TypeOf(mapdemo["fuhongxue"]).String())
	fmt.Println(str)
	// for k, v := range mapdemo {
	// 	if reflect.TypeOf(v).String() == "string" {
	// 		fmt.Println("value type is string , val = ", v.(string))
	// 	}
	// }
}
