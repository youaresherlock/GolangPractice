package main

import (
	"fmt"
	"encoding/json"
	"log"
)

/*
golang json tag omitempty序列化时会忽略空值
json:"-" 表示不进行序列化
*/

// json序列化和反序列化
type Student struct {
	ID int 
	Name string 
	EnglishName string `json:"clarence,omitempty"`
	BlogUrl string `json:"-"`
}

func main() {
	s := Student {
		ID: 1,
		Name: "狗蛋",
		EnglishName: "", 
		BlogUrl: "http://blog.csdn.net/qq_122344",
	}
	fmt.Println(s)
	// 序列化
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("marshal error:%s", err)
	}
	fmt.Println(string(buf))

	str := `{"ID": 2, "Name": "clarence"}`
	var s2 Student 
	err1 := json.Unmarshal([]byte(str), &s2)
	if err1 != nil {
		log.Fatalf("unmarshall error:%s", err)
	}
	fmt.Println(s2)
}