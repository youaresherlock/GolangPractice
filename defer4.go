/*
* @Author: Clarence
* @Date:   2019-09-12 09:25:37
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-12 09:49:04
*/

/*
defer, return和返回值之间的执行顺序 
defer,返回值和return三者的执行顺序是: 
return最先给返回值赋值, 接着defer开始执行一些收尾工作, 
最后RET指令携带返回值退出函数
*/

package main 

import "fmt"

func main() {
	fmt.Println("return: ", a())
	fmt.Println("return: ", b())

}

// 无名返回值
func a() int {
	var i int 
	defer func() {
		i++ 
		fmt.Println("defer2: ", i, &i)
	}()
	defer func() {
		i++
		fmt.Println("defer1: ", i, &i)
	}()
	//  返回值没有被声明,因此返回的还是0 
	return i
}

// 有名返回值
func b() (i int) {
	defer func() {
		i++ 
		fmt.Println("defer2: ", i, &i)
	}()
	defer func() {
		i++
		fmt.Println("defer1: ", i, &i)
	}()
	/*
	b() (i int)已经声明了i, defer中是可以调用到真实的返回值的, 
	defer在return赋值返回值i之后, 再一次修改了i的值, 因此返回2
	*/
	return i
}

// defer1:  1 0xc000014100
// defer2:  2 0xc000014100
// return:  0
// defer1:  1 0xc000014140
// defer2:  2 0xc000014140
// return:  2