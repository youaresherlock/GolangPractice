package main
import (
	"fmt"
	// "time"
	// "reflect"
)

func main() {
	testDerive()
}

type A struct {
	aa int
	BB string
}

type B struct {
	A
	aa int
	CC string
}

func (a *A) aFunc() {
	fmt.Println("A.aFunc")
}

func (a *A) BFunc() {
	fmt.Println("A.BFunc")
}

func (b *B) aFunc(test1, test2 int, str ...string) {
	fmt.Println("B.aFunc")
	fmt.Println(test1, test2, str)
}

func (b *B) CFunc() {
	fmt.Println("B.CFunc")
}

func testDerive() {
	// var tb B
	// //变量测试
	// fmt.Println(tb.aa)
	// fmt.Println(tb.A.aa)
	// fmt.Println(tb.BB)
	// fmt.Println(tb.A.BB)
	// fmt.Println(tb.CC)
	// tb.aa = 9
	// tb.A.aa = 8
	// tb.BB = "A.BB"
	// tb.CC = "B.CC"
	// fmt.Println(tb.aa)
	// fmt.Println(tb.A.aa)
	// fmt.Println(tb.BB)
	// fmt.Println(tb.A.BB)
	// fmt.Println(tb.CC)

	// //函数测试
	// tb.aFunc()
	// tb.A.aFunc()
	// tb.BFunc()
	// tb.A.BFunc()
	// tb.CFunc()

// 空interface不包含任何的method,所有的类型都实现了空interface
// 函数可以将interface{}作为参数，也可以作为返回值

	// var inter interface{}
	// inter = tb
	// fmt.Println(inter.(B).aa)

	// 可变参数可以将一个或者多个参数赋值个这个占位符，也可以不传入,其他指定的参数不许传入
	//0001-01-01 00:00:00 +0000 UTC
	// var t time.Time
	// t = time.Now()
	// if ! t.IsZero() {
	// 	fmt.Println(t)
	// }
	// t = time.Time{}
	// fmt.Println(t.IsZero())

	// test := []A{}
	// test1 := []*A{}
	// fmt.Println("test type: ", reflect.TypeOf(test))
	// fmt.Println("test1 type: ", reflect.TypeOf(test1))

// zero是整形，不是布尔型， 不能直接使用if判断和python, c, c++不同
	// zero := 0
	// if zero == 0{
	// 	fmt.Println(zero)
	// 	fmt.Println(time.Now().Format("2006-01-02-15:04:05"))
	// }

	//高精度类型不能赋值给低精度类型, go不同类型之间不能赋值，不像c强类型语言
	var a float64
	a = 1
	fmt.Println("++++++++++++++++", a) //此处可以赋值
	// a = 1
	// fmt.Println("+++++++++++++++++", a)
	// var b float32
	// b = a
	// fmt.Println("================", b)
	// var t time.Time
	// t = time.Now()
	// fmt.Println(t)
	if true == 0 {
		fmt.Println("true = 0")
	}
}
