package main 

import "fmt"

func main() {
	// 传递int类型多个参数
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{1, 2, 9, 8, 8}
	// 直接传递slice切片
	x = min(slice...)
	fmt.Printf("The minimum is: %d\n", x)
}

// 如果函数的最后一个参数是...type的形式，就可以处理可变长的参数，长度可以为0
func min(s ...int) int {
	if len(s) == 0{
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}