/*
* @Author: Clarence
* @Date:   2019-07-13 19:21:45
* @Last Modified by:   Clarence
* @Last Modified time: 2019-07-13 19:36:50
*/
/*
类型断言:
接口可以包含任何类型的值，类型断言是检测接口运行时在变量中存储的值的实际类型v := varI.(T)
更安全的类型断言方式:
if v, ok := varI.(T); ok {

}
断言成功则v是转换到类型T的值,ok是true;否则v是类型T的零值, ok是false
*/
package main

import "fmt"


/*contract that defines different things that have value*/
type Valuable interface {
	getValue() float32
}

type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	maker string 
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

func showValue(asset Valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var o Valuable = stockPosition{"GOOD", 10.5, 4}
	showValue(o)

	o = car{"BMW", "M3", 66500}
	showValue(o)

	switch  t := o.(type) {
		case stockPosition:
			fmt.Println("stockPosition")
		case car:
			fmt.Println("car", t)
	}
}