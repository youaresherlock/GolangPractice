package main 

import (
	"fmt"
	"strconv"
)

func main() {
	var temp int = 100
	fmt.Println(string(temp))
	fmt.Println(strconv.Itoa(temp))
	fmt.Println(strconv.FormatInt(int64(temp), 10))

//接口变量的类型可以使用type-switch检测 类型断言 
//类型强制转换是发生在编译期间，类型断言在运行期间才能确定
	var tmp interface{} = temp
	switch tmp.(type) {
	case int: 
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}