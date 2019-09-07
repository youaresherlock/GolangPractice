package main
import (
   "fmt"
)
type Student struct {
   name string
   age  int
}
type Class struct {
   name string
   age  int
}
type School struct {
   Student
   Class
}
func main() {
   var t School
   t.Student.name = "stu"
   t.Class.name = "cls"
   fmt.Println(t)
}