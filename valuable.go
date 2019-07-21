/*
* @Author: Clarence
* @Date:   2019-07-13 19:02:43
* @Last Modified by:   Clarence
* @Last Modified time: 2019-07-13 19:10:58
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
}