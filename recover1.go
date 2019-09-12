/*
* @Author: Clarence
* @Date:   2019-09-12 10:18:17
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-12 10:23:35
*/
package main 

import (
	"fmt"
	"time"
)

func main() {
	throwsPanic(genErr) 
}

// 产生错误函数 
func genErr() {
	fmt.Println(time.Now(), "正常的语句")
	defer func() {
		fmt.Println(time.Now(), "defer 正常的语句")
		panic("第二个错误")
	}()

	panic("第一个错误")	//不影响defer的执行 
}

// 抛出异常并使用panic恢复
func throwsPanic(f func()) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(time.Now(), "捕获到的异常: ", r)
			b = true 
		}
	}()

	f() 

	return 
}


// 2019-09-12 10:23:22.6383794 +0800 CST m=+0.015931101 正常的语句
// 2019-09-12 10:23:22.7381675 +0800 CST m=+0.115719201 defer 正常的语句
// 2019-09-12 10:23:22.7391429 +0800 CST m=+0.116694601 捕获到的异常:  第二个错误