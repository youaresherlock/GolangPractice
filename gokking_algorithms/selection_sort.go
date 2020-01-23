package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)
	for i := 1; i < size; i++ {
		smallestIndex := findSmallest(arr)
		newArr[i] = arr[smallestIndex]
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return newArr
}

func main() {
	s := []int{1, 3, 10, 6, 19}
	fmt.Println(selectionSort(s))
}
