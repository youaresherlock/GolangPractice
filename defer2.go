package main 

import "fmt"

func trace(s string) {fmt.Println("entering:", s)}
func untrace(s string) {fmt.Println("leaving:", s)}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func main() {
	b()
}

/*
使用defer实现代码追踪
entering: b
in b
entering: a
in a
leaving: a
leaving: b
*/



