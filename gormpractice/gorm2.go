/*connect mysql*/
package main

import (
	"fmt"
	// "time"
	// "database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  )

  type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
  }

  func (Animal) TableName() string {
	  return "allanimals"
  }
  
  func main() {
	db, err := gorm.Open("mysql", "root:x1430371727@/news_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	db.AutoMigrate(&Animal{})

	var animal = Animal{Age: 99, Name: "clarence"}
	db.Create(&animal)
	
	fmt.Println(animal.ID, animal.Name, animal.Age)
  }


/*
Token: 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjE2NDQxMzIsImp0aSI6MjU1LCJieGlkIjoyNTUsImNvbXBhbnlfaWQiOjEsImlzcyI6Imh0dHA6Ly90YXBpLmJvc3MuZnVueGRhdGEuY29tLyIsInN1YiI6ImVtcGxveWVlIn0.skNy-vEkTDCwXyjvH_hZXxv0cGNg8_DQ2A9qTbFLkBE
*/