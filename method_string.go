/*
String()方法来定制类型的字符串形式的输出--一种可阅读性和打印性的输出
*/
package main 

import (
	"fmt"
	"strconv"
	"runtime"
)

type TwoInts struct {
	a int 
	b int 
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)

	// 已分配内存的总量 runtime包访问GC进程
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}

func (tn *TwoInts)String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")" 
}