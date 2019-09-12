/*
* @Author: Clarence
* @Date:   2019-09-12 10:00:24
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-12 10:12:42
*/
package main 

import "fmt"

func main() {
	b() 
}

func a() {
	defer un(trace("a"))
	fmt.Println("a的逻辑代码")
}

func b() {
	defer un(trace("b"))
	fmt.Println("b的逻辑代码")
	a() 
}

func trace(s string) string {
	fmt.Println("开始执行", s)
	return s 
}

func un(s string) {
	fmt.Println("结束执行", s)
}

// 开始执行 b
// b的逻辑代码
// 开始执行 a
// a的逻辑代码
// 结束执行 a
// 结束执行 b