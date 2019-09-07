package main 


/*
是否可设置是Value的一个属性,不是所有的反射值都有这个属性
CanSet() 
当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，
那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)。
*/
import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4 
	v := reflect.ValueOf(x) //传递x的拷贝创建v
	// setting a value:
    // v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
    fmt.Println("settability of v:", v.CanSet())
    v = reflect.ValueOf(&x) // Note: take the address of it 
    fmt.Println("type of v:", v.Type())
    fmt.Println("settability of v:", v.CanSet())
    v = v.Elem()
    fmt.Println("The Elem of v is: ", v)
    fmt.Println("settability of v:", v.CanSet())
    v.SetFloat(3.1415)
    fmt.Println(v.Interface())
    fmt.Println(v)
}