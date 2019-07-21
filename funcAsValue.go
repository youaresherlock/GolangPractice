/*Go语言提供了灵活性，可以即时创建函数并将其用作值使用.*/
package main

import (
	"fmt"
	"math"
)

func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}