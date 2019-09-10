/*
* @Author: Clarence
* @Date:   2019-09-10 21:48:59
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-10 21:52:24
*/
package main 

import "fmt"
 

 func main() {
 	LOOP1: 
 		for {
 			x := 1
 			switch {
 			case x > 0: 
 				fmt.Println("A")
 				// 跳出到标签为LOOP1的代码块之外
 				break LOOP1
 			case x == 1:
 				fmt.Println("B")
 			default: 
 				fmt.Println("C")
  			}
 		}
 }