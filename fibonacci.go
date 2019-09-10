/*
* @Author: Clarence
* @Date:   2019-09-10 22:10:19
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-10 22:12:46
*/
package main 

import "fmt"
 

// 0 1 1 2 3 5 8 ... 
 func fibonacci(n int) int {
 	if n < 2 {
 		return n
 	}

 	return fibonacci(n - 1) + fibonacci(n - 2)
 }

 func main() {
 	var i int 
 	for i = 0; i < 10; i ++ {
 		fmt.Printf("%d\t", fibonacci(i))
 	}
 }

 // 0       1       1       2       3       5       8       13      21      34