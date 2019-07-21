/*go提供了接口(interfaces)的数据类型，它代表一组方法签名
struct数据类型实现这些接口以具有接口的方法签名的方法定义

type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}
type struct_name struct {

}

func (struct_name_variable struct_name) method_name1() [return_type] {
//    method implementation 
}
多态: 是面向对象编程中一个广为人知的概念. 根据当前的类型选择正确的方法,或者说: 同一种类型
在不同的实例上似乎表现出不同的行为
*/

package main

import (
	"fmt"
	"math"
	"reflect"
)

/*define a interface*/
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

/*define a method for circle (implementation of Shape)*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius 
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x:0, y:0, radius:5}
	rectangle := Rectangle{width:10, height:5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
	fmt.Println(reflect.TypeOf(circle))
}