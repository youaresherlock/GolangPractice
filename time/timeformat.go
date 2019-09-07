package main

import (
	"fmt"
	"time"
	// "unsafe"
	// "reflect"
)

type Str string 

type Test struct {
	Name string 
}

func main() {
	// now := time.Now()

	// fmt.Println(now)

	// // go时间格式化必须精确限定到golang指定的时间原点	2006-1-2 15:04:05 golang2016内部谋划, 1,2,3,4,5顺序表示
	// fmt.Println(now.Format("2006-01-02 15:04:05"))
	// fmt.Println(now.Format("2006-1-2 15:04:05"))
	// fmt.Println(now.Format("2006/1/2 25:04:05")) // 2019/7/10 1022:32:22

	// //获取时间戳
	// timestamp := time.Now().Unix()
	// fmt.Println(timestamp)

	// // 时间戳转化为字符串时间
	// fmt.Println(time.Unix(timestamp, 0).Format("2006-01-02 15:04:05"))

	// //从字符串时间转换为时间戳
	// tm2, _ := time.Parse("01/02/2006 15:04:05", "02/08/2015 3:05:56") //2019-07-10 11:02:44

	// fmt.Println(tm2.Unix())

	// tm3 := time.Date(2019, 7, 1, 0, 0, 0, 0, time.Local)
	// fmt.Println(tm3)

	// var str Str 
	// str = "testest"
	// fmt.Println(str)

	// test1 := &Test{}
	// // test2 := &Test{}
	// fmt.Println(unsafe.Sizeof(*test1))
	// fmt.Println(test1.Name == "")
	// var test float64
	// fmt.Println(test)

	requestTime, err := time.Parse("2006-01-02", "2019-06-09")
	if err != nil {
		panic("时间格式错误") 
	}
	fmt.Println(requestTime)

	var test = int(time.Date(2019, 8, 19, 0, 0, 0, 1, time.Local).Weekday())


	a, b := 0, 7
	fmt.Println(test, a | b)
}