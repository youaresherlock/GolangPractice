/*
方法是使用接收器的特殊函数
Go语言中一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值
或者是一个指针.所有给定类型的方法属于该类型的方法集
func (variable_name variable_data_type) function_name() [return_type] {
	//  函数体
}
*/
package main

import (
	"fmt"
	"math"
)

/*define a circle*/
type Circle struct {
	x, y, radius float64
}

/*define a method for circle*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	circle := Circle{x:0, y:0, radius:5}
	fmt.Printf("Circle area: %f", circle.area())
}

