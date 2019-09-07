package main

import (
	"fmt"
)

func main() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	var slice []byte;
	// 切片可以用数组来初始化，是数组的引用，数组改变，切片会改变;也可以通过内置函数make()初始化
	slice = ar[2:4]
	fmt.Println("初始的slice ", slice, cap(slice))
	ar[2] = 'h'
	ar[3] = 'i'
	fmt.Println("以后的slice ", slice)
	slice2 := make([]byte, 10, 20)
	fmt.Println(len(slice2), cap(slice2))
}