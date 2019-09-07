package main

import "fmt"

func main() {
	/*简短声明的变量只能在函数内部使用
	不能用简短声明方式来单独为一个变量重复声明, :=左侧至少有一个新变量，才允许多变量的重复声明
	struct 的变量字段不能使用 :=来赋值*/
	one := 0
	one, two := 1, 2
	one, two = two, one
	fmt.Println(one, two)
}