package main

import (
	"fmt"
)

func main() {
	/*
	new和make均是用于分配内存: new用于值类型和用户定义的类型, 如自定义结构,make用于内置引用类型(切片/map/管道)
	将类型作为参数，分配类型的零值并返回其地址
	*/
	temp1 := new(int)
	fmt.Println("temp1", temp1, *temp1) //temp1 0xc000056058 0
	a := []int{1,2}[0]
	fmt.Println(a)

	test := map[string]string{"hello" : "world", "name" : "clarence"}
	for _, item := range test{
		fmt.Println(item)
	}
	for _, item := range test{
		fmt.Println(item)
	}
}

//未生效和已经生效