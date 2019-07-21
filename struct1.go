// 结构体字段使用点号来访问
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// 结构体字段可以通过结构体指针来访问
	p.X = 4
	fmt.Println(v, p.X)
}