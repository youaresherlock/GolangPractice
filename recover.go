/*
* @Author: Clarence
* @Date:   2019-09-12 10:12:54
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-12 10:18:08
*/

// recover可以捕获到panic的输入值,并且恢复正常的执行 
package main 

import "log"

func main() {
	test() 
}

func test() {
	defer func() {
		//  只会捕获最后一个错误
		if r:= recover(); r != nil {
			log.Printf("捕获到的异常: %v", r)
		}
	}()

	defer func() {
		panic("第二个错误")
	}()

	panic("第一个错误")	//不影响defer的执行 
}