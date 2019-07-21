/*
* @Author: Clarence
* @Date:   2019-06-28 18:56:05
* @Last Modified by:   Clarence
* @Last Modified time: 2019-06-28 19:04:13
*/
// 结构体继承
package main
import (
	"fmt"
)

func main() {
	testDerive()
}

type A struct {
	aa int
	BB string
}

type B struct {
	A
	aa int
	CC string
}

func (a *A) aFunc() {
	fmt.Println("A.aFunc")
}

func (a *A) BFunc() {
	fmt.Println("A.BFunc")
}

func (b *B) aFunc() {
	fmt.Println("B.aFunc")
}

func (b *B) CFunc() {
	fmt.Println("B.CFunc")
}

func testDerive() {
	var tb B
	//变量测试
	fmt.Println(tb.aa)
	fmt.Println(tb.A.aa)
	fmt.Println(tb.BB)
	fmt.Println(tb.A.BB)
	fmt.Println(tb.CC)
	tb.aa = 9
	tb.A.aa = 8
	tb.BB = "A.BB"
	tb.CC = "B.CC"
	fmt.Println(tb.aa)
	fmt.Println(tb.A.aa)
	fmt.Println(tb.BB)
	fmt.Println(tb.A.BB)
	fmt.Println(tb.CC)

	//函数测试
	tb.aFunc()
	tb.A.aFunc()
	tb.BFunc()
	tb.A.BFunc()
	tb.CFunc()

	test := &B{aa: 10, CC: "test"}
}