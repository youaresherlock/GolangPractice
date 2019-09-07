package main

import (
	"fmt"
)

func main() {
	// 1 声明类型,没有初始化,零值(false)
	var a bool
	a = true
	fmt.Println("a = ", a)
	 
	 
	// 2 自动推导类型
	var b = false
	fmt.Println("b = ", b)
	 
	 
	c := true
	if c {
		fmt.Println("c = ", c)	
	}
}
租户ID 和 用户ID-微信号生成