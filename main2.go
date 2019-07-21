/*函数可以没有参数或接受多个参数*/
package main 

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}