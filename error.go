/*
go编程提供了一个非常简单的错误处理框架，以及内置的错误接口类型
type error interface {
	Erorr() string
}
函数通常返回错误作为最后一个返回值,可使用errors.New来构造一个基本的错误消息
func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt)
	}
	return math.Sqrt(value)
}
result, err := Sqrt(-1)

if err != nil {
	fmt.Println(err)
}
*/
package main

import "errors"
import "fmt"
import "math"

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
