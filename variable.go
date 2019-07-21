// package main

// import "fmt"

// func main() {
// 	/*在函数中, :=简洁赋值语句在明确类型的地方，可以用于替代var定义
// 	:=结构不能使用在函数外*/
// 	var i, j int = 1, 2
// 	k := 3
// 	c, python, java := true, false, "no"
// 	fmt.Println(i, j, k, c, python, java)
// }

package main

import "fmt"

const (
    Big   = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}