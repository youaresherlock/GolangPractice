/*
* @Author: Clarence
* @Date:   2019-09-10 21:52:57
* @Last Modified by:   Clarence
* @Last Modified time: 2019-09-10 22:04:00
*/

/* 
函数类型: 所有拥有相同参数与相同返回值的一种函数类型
type typeName func(intput1 intputType1, intput2 inputType2[, ...]) (result1, resultType1 [, ...])
*/
package main 

import "fmt"

func isOdd(v int) bool {
	if v % 2 == 0 {
		return false 
	}

	return true 
}

func isEven(v int) bool {
	if v % 2 == 0 {
		return true 
	}

	return false 
}


type boolFunc func(int) bool // 声明一个函数类型 

func filter(slice []int, f boolFunc) []int {
	var result []int 
	for _, value := range slice {
		if f(value) {
			result = append(result, value) 
		}
	}
	return result 
}

func main() {
	slice := []int{3, 1, 4, 5, 9, 2}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd: ", odd) 
	even := filter(slice, isEven)
	fmt.Println("even:", even)
}


// slice =  [3 1 4 5 9 2]
// odd:  [3 1 5 9]
// even: [4 2]
