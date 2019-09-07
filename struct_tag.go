/*
reflect包实现了运行时反射，允许程序操作任意类型的对象
结构体中的字段除了有名字和类型外,还可以有可选的标签(tag)
reflect.TypeOf()可以获取变量的正确类型
*/
package main

import (
    "fmt"
    "reflect"
)


type TagType struct { // tags
    field1 bool   "An important answer"
    field2 string "The name of the thing"
    field3 int    "How much there are"
}

func main() {
    tt := TagType{true, "Barak Obama", 1}
    for i := 0; i < 3; i++ {
        refTag(tt, i)
    }
}

func refTag(tt TagType, ix int) {
    ttType := reflect.TypeOf(tt)
    ixField := ttType.Field(ix)
    value := reflect.ValueOf(tt)
    fmt.Println(value, ixField.Tag)
}