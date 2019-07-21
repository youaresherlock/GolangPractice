// for循环的range格式可以对slice或map进行迭代循环 
// 每次迭代range返回两个值,第一个是当前下标，第二个是该下标所对应元素的一个拷贝
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

