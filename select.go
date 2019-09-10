/*
* @Author: Clarence
* @Date:   2019-09-09 22:36:56
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-09 22:44:03
*/
package main 

import "fmt"

func main() {
	a := make(chan int, 1024)
	b := make(chan int, 1024)

	for i := 0; i < 10; i ++ {
		fmt.Printf("第%d次,", i)

		a <- 1
		b <- 1

		select {
		case <-a:
			fmt.Println("from a")
		case <-b:
			fmt.Println("from b")
		}
	}
}

// 第0次,from b
// 第1次,from a
// 第2次,from a
// 第3次,from b
// 第4次,from b
// 第5次,from b
// 第6次,from b
// 第7次,from a
// 第8次,from b
// 第9次,from b