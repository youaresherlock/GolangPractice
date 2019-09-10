/*
* @Author: Clarence
* @Date:   2019-09-09 22:06:59
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-09 22:20:11
*/
package main 

import "fmt"

func main() {
	sum := 1
	count := 3
	mean := float32(sum) / float32(count)
	mean2 := sum / count
	// mean2 := sum / float(count) invalid operation: sum / float32(count) (mismatched types int and float32)
	fmt.Println("mean的值为: %f\n", mean)
	fmt.Println("mean2的值为: %d\n", mean2)

	// nil 
	var ptr *int 
	aP := &ptr
	fmt.Println("ptr的值是: %x\n", ptr) // <nil>
	fmt.Println("aP的值是: %x\n", aP) // 0xc000092020


}