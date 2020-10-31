package main

import (
	"fmt"
	"math"
)

//例子：求平面直角坐标系内两个点的距离
type Point struct {
	x, y float64
}

func getDistance(p1, p2 Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

//站在点的角度考虑问题，计算点之间的距离,计算我到某个点的距离
func (this Point) getDis(p Point) float64 {
	return math.Sqrt((this.x-p.x)*(this.x-p.x) + (this.y-p.y)*(this.y-p.y))
}
func (p Point) show() {
	fmt.Println("x = ", p.x, ",y = ", p.y)
}

func main() {
	p1 := Point{0.0, 0.0}
	p2 := Point{3.0, 4.0}
	fmt.Println(getDistance(p1, p2))
	fmt.Println(p1.getDis(p2))
	fmt.Println(p2.getDis(p1))
	p2.show()
}
