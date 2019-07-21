// 切片可以包含任意的类型
package main


import (
	"fmt"
	"strings"
)

func main() {
	// create a tic-tac-toe board
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		// func Join(a []string, sep string) string 将一系列字符串连接为一个字符串，之间用sep来分隔
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}