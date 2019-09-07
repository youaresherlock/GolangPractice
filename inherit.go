package main
import "fmt"

/*
   如果一个struct 嵌套了另一个匿名结构体,那么这个结构可以直接访问匿名结构体的方法,从而实现了继承
   如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现了继承。

​   如果一个struct嵌套了另一个有名结构体，那么这个模式就叫组合。

   多重继承：如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现了多重继承。
*/

type Student struct {
   Id int
   name   string
}
func (p *Student) Run() {
   fmt.Println("greg")
}
type Class struct {
   Student
   golang string
}
type School struct {
   gregs Student
}
func main() {
   var c Class
   c.Id = 100
   c.name = "greg2"
   c.golang = "一班"
   fmt.Println(c)
   c.Run()
   var sch School
   sch.gregs.Id = 100
   sch.gregs.name = "train"
   fmt.Println(sch)
   sch.gregs.Run()
}