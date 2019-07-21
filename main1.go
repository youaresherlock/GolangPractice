package main 

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	/*首字母大写的名称是被导出的,任何未导出的名字是不能被包外的代码访问的*/
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("My favorite number is", math.Pi)
}