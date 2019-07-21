/*在程序中，局部变量和全局变量的名称可以相同,但函数内部的局部变量的值将优先
(即局部变量会覆盖全局变量)
*/
package main

import "fmt"

/*global variable declaration*/
var g int = 20

func main() {
	/*local variable declaration*/
	var g int = 10

	fmt.Printf("value of g = %d\n", g)
}