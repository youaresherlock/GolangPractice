package main

import (
   "fmt"
)

type user struct {
   name string 
   age byte 
}


func main() {
   user := user{"Tom", 2} // 按声明时的字段顺序初始化

   /*
   匿名结构体,与其他定义类型的变量一样,如果在函数外部需在结
   构体变量前加上var关键字，在函数内部可省略var关键字
    var config struct {
    APIKey string
    OAuthConfig oauth.Config
   }
   */

   //定义并初始化并赋值给data
   data := struct {
      Title string 
      Content string 
   }{
      "tensorflow实战",
      "tensorflow和pytorch",
   }
   
   fmt.Println(user, data) // {Tom 2} {tensorflow实战 tensorflow和pytorch}
}