package main  
import "fmt"
var (     
    a int = b + 1     
    b int = 1

)  
func main() {     
    fmt.Println(a)     
    fmt.Println(b) 
    // 函数作用域内的局部变量，初始化顺序是从左到右,从上到下
    var (
        c int = d + 1
        d int = 1
    )
    fmt.Println(c)
    fmt.Println(d)
}