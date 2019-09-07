/*
Golang支持两种类型的字符串字面量
解释形字符串 " "
非解释形字符串 ` ` 转移字符不会被解释，并且还支持换行

和 C/C++ 不一样，Golang 语言中的字符串是根据长度限定的，
而非特殊的字符 \0。string 类型的 0 值是长度为 0 的字符串，即空字符串 ""

golang中的byte类型和rune类型
byte等同于int8 常用来处理ASCII字符
rune等同于int32 常用来处理unicode或utf-8字符
*/
package main 

import "fmt"

func main() {
	s1 := "Hello\nWorld"
	s2 := `Hello\n
	nick!`
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := "abc你"
	// 内置函数返回一个字符串中的字节数, 而不是字符数
	fmt.Printf("字符串的字节长度是: %d\n", len(s3)) // 6
	for i := 0; i < len(s3); i ++ {
		fmt.Println(s3[i])
	}

	// 获取字符串中的字符数 type rune =  int32
	r := []rune(s3)
	fmt.Println(r, len(r)) // [97 98 99 20320] 4

	// 在字符串中含有非单字节的字符时这种方法是不正确的.range函数能解决这个问题
	for _, v := range s3 {
		fmt.Printf("%c", v)
	}


	// 修改字符串  修改字节 []byte /  修改字符 []rune

	// 修改字符串中的字节 
	s4 := "Hello 世界"
	b := []byte(s4)
	b[5] = '.'
	fmt.Printf("%s\n", s4)
	fmt.Printf("%s\n", b)

	// 修改字符串中的字符 
	s5 := "Hello 世界"
	c := []rune(s5)
	c[6] = '中'
	c[7] = '囯'
	fmt.Println(s5, string(c))
}

