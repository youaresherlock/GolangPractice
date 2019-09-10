/*
* @Author: Clarence
* @Date:   2019-09-10 22:05:42
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-10 22:07:14
*/

// 匿名函数是指不需要定义函数名的一种函数定义方式

package main 

import "fmt"

func main() {
	func(num int) int {
		sum := 0 
		for i := 1; i <= num; i++ {
			sum += i 
		}
		fmt.Println(sum)
		return sum 
	}(100)
}