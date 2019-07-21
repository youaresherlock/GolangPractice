/*
* @Author: Clarence
* @Date:   2019-07-13 19:57:15
* @Last Modified by:   Clarence
* @Last Modified time: 2019-07-13 20:00:11
*/
/*
方法是函数，不允许方法重载,即对于一个类型只能有一个给定名称的方法.但是
如果基于接受者类型，是有重载的
*/
package main

import "fmt"

type TwoInts struct {
    a int
    b int
}

func main() {
    two1 := new(TwoInts)
    two1.a = 12
    two1.b = 10

    fmt.Printf("The sum is: %d\n", two1.AddThem())
    fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

    two2 := TwoInts{3, 4}
    fmt.Printf("The sum is: %d\n", two2.AddThem())
}

func (tn TwoInts) AddThem() int {
    return tn.a + tn.b
}

func (tn TwoInts) AddToParam(param int) int {
    return tn.a + tn.b + param
}