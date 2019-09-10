/*
* @Author: Clarence
* @Date:   2019-09-10 21:36:12
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-10 21:46:45
*/
/*
defer用于延迟调用指定函数， defer关键字只能出现在函数内部
1. 只有当defer语句全部执行, defer所在函数才算真正结束执行 
2. 当函数中有defer语句时, 需要等待所有defer语句执行完毕, 才会执行return语句
defer语句用于回收资源,清理收尾等工作
*/
// package main 

// import "fmt"

// var i = 0

// func print() {
// 	fmt.Println(i)
// }

// func main() {
// 	for ; i < 5; i++ {
// 		defer print()
// 	}
// }

// 5 
// 5 
// 5 
// 5 
// 5 

package main 

import "fmt"

var i = 0

func print(i int) {
	fmt.Println(i)
}

func main() {
	for ; i < 5; i++ {
		// 开始压入栈中, 先入后出
		defer print(i)
	}
}

// 4
// 3
// 2
// 1
// 0