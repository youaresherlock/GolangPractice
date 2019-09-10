package main

import (
	"fmt"
	"runtime"
)

var x interface{} // 空接口

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Printf("OS X.")
	case "linux":
		fmt.Printf("Linux.")
	default: 
		fmt.Printf("%s.", os)
	}

	x = 1
	switch i := x.(type) {
	case nil:
		fmt.Printf("nil, x的类型是%T", i)
	case int:
		fmt.Printf("int, x的类型是%T", i)
	default:
		fmt.Printf("未知类型")
	}

	switch marks := 90; {
	case marks > 80:
		fmt.Println(">80")
	}
}