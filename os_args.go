// os包中有一个string类型的切片变量os.Args, 用来处理一些基本的命令行
// 参数, 它在程序启动后读取命令行输入的参数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		// func Join(a []string, sep string) string) 第一个参数为程序执行路径
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
