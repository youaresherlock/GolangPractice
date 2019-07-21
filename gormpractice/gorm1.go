/*connect mysql*/
package main

import (
	"fmt"
	"time"
	// "database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  )

  type User struct {
	  gorm.Model
	  Name string `gorm:"size:20;not null"`
	  Age int
	  Birthday time.Time
	  Email string `gorm:"type:varchar(100);unique_index"`
	  Address string `gorm:"index:addr"`
	  School string 
	  IgnoreMe int `gorm:"-"`
  }

  func (User) TableName() string {
	  return "profiles"
  }
  
  func main() {
	db, err := gorm.Open("mysql", "root:x1430371727@/news_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// db.AutoMigrate(&User{})

	// db.Create(&User{
	// 	Name: "clarence", 
	// 	Age: 24,
	// 	Birthday: time.Now(),
	// 	Email: "1234@qq.com",
	// 	Address: "xi'an",
	// 	School: "unknown"})

	// var user User
	// db.First(&user, "email = ?", "1234@qq.com") 

	// fmt.Println(user.Name, user.Email, user.Address)

	// var user1 User
	// user1 = User{Name: "John", Age: 20, Birthday: time.Now(), Email: "12345@qq.com", Address: "tianjin", School: "university of technology"}
	// db.Save(&user1)
	// db.Model(&user).Update("Address", "beijing")

	// var user2 User
	// db.First(&user2, "email = ?", "12345@qq.com")
	// fmt.Println(user2.Name, user2.Email, user2.Address)
	// db.Delete(&user)

	// var users []User

	// db.Where("name = ?", "Bob").Find(&users)
	// db.Where("name = ?", "clarence").Or("email = ?", "1234@qq.com").Find(&users)
	// db.Where(User{Name: "clarence"}).Or(User{Email: "1234@qq.com"}).Find(&users)
	// db.Where(map[string] interface{}{Name: "clarence"}).Or(map[string] interface{}{Email: "1234@qq.com"}).Find(&users)
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name, user.Email, user.Address)
	// }

	// Get first matched record, or initalize a new one with given conditions (only works with struct, map conditions
	// var user User
	// db.Where(User{Name: "Jinzhu"}).FirstOrInit(&user)
	// fmt.Println(user.Name)


	// 更新单个属性
	var user User
	db.Where("name = ?", "clarence").First(&user)
	db.Model(&user).Update("address", "USA")
	fmt.Println(user.Name, user.Address)
  }

//  go run main.go serve --addr :8888 
// go run main.go serve --addr :8888 -D 进入调试模式 执行的sql语句会有日志
// https://golang.org/x/tools/cmd/guru