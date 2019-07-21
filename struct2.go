package main

import "fmt"

type Vertex struct{
	X, Y int
}

var (
	v1 = Vertex{4, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{4, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}