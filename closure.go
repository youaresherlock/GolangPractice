/*
go语言支持可以充当函数闭包的匿名函数.当想要定义一个函数内联而不传递任何名称时,
使用匿名函数
*/
package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	
	/*create a new sequence and see the result, i is 0 again*/
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

}

// result: 1 2 3 1 2

/* 
// 闭包带参数
package main
import "fmt"
func main() {
    add_func := add(1,2)
    fmt.Println(add_func(1,1))
    fmt.Println(add_func(0,0))
    fmt.Println(add_func(2,2))
} 
// 闭包使用方法
func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
    i := 0
    return func(x3 int,x4 int) (int,int,int){ 
       i++
       return i,x1+x2,x3+x4
    }
}
*/