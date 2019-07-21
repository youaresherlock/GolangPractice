/*
strings包含一个用于连接多个字符串的join()方法
strings.Join(sample, "")
Join连接数组的元素以创建单个字符串。第二个参数是分隔符，放置在数组的元素之间。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := []string{"hello", "world"}
	fmt.Println(strings.Join(greetings, " "), greetings[0])
}