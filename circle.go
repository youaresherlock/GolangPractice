// Go只有一种循环结构-for循环
package main 

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	/*
	可以省略分号为while循环
	for sum < 1000 {
		sum += sum
	}
	死循环
	for {
		
	}
	*/
	fmt.Println(sum)
}