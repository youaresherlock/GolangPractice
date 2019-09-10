/*
* @Author: Clarence
* @Date:   2019-09-10 21:29:18
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-10 21:34:47
*/
package main 

import "fmt"

// 将NewInt定义为int类型
type NewInt int 
// 将int取一个别名叫IntAlias 
type IntAlias = int 

func main() {
	var a NewInt 
	fmt.Printf("a type: %T\n", a)
	var a2 IntAlias 
	fmt.Printf("a2 type: %T\n", a2)
}
/*
a的类型是main.NewInt,表示main包下定义的newInt类型.
a2类型是int. IntAlias类型只会在代码中存在, 编译完成时不会有Alias类型
*/
// a type: main.NewInt
// a2 type: int