package main

import (
	"fmt"
)

type IntArray []int

func (p IntArray) Print() {fmt.Println(p)}

func main() {
	p := []int{1,2,3,4}

	IntArray(p).Print()
}

