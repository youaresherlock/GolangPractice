/*
反射使用程序检查其所拥有的结构,尤其是类型的一种能力.
反射可以在运行时检查类型和变量,例如它的大小、方法和
动态的调用这些方法
reflect.Type reflect.Value有许多方法用于检查和操作它们
Value有一个Type方法返回reflect.Value的Type 
Type和Value都有Kind方法返回一个常量来表示类型 Uint Float64 Slice....
Value有Int,Float的方法可以获取存储在内部的值
const (
    Invalid Kind = iota
    Bool
    Int
    Int8
    Int16
    Int32
    Int64
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
    Float32
    Float64
    Complex64
    Complex128
    Array
    Chan
    Func
    Interface
    Map
    Ptr
    Slice
    String
    Struct
    UnsafePointer
)
*/
package main 

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Float())
	fmt.Printf("kind: %d\n", v.Kind()) // 14
	fmt.Println("kind:", v.Kind()) // float64
 	fmt.Printf(v.Kind().String()) // float64
	fmt.Println(v.Interface()) //得到还原(接口)值
	y := v.Interface().(float64)
	fmt.Println(y)
}


/* output:
type: float64
value: 3.4
type: float64
value: 3.4
kind: 14
kind: float64
float643.4
3.4
*/