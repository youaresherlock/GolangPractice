package main 

import (
	"fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
  gorm.Model
  Code string
  Price uint 
}

type Output struct {
	Name string
	Price uint 	`gorm:"column:prices"`
	Code string
}

func main() {
  db, err := gorm.Open("mysql", "root:x1430371727@/test?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("连接数据库失败")
  }
  defer db.Close()

  // 自动迁移模式
  // db.AutoMigrate(&Product{})

  // 创建
  // db.Create(&Product{Code: "L1212", Price: 1000})

  // // 读取
  // var product Product
  // db.First(&product, 1) // 查询id为1的product
  // db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

  // 更新 - 更新product的price为2000
  // db.Model(&product).Update("Price", 2000)

  // // 删除 - 删除product
  // db.Delete(&product)
  // Scan映射规则: 根据查询出来的字段名字映射到结构体中的字段名字 
  // 无select会自动映射，有会限制映射的字段
  var out Output
  // db.Table("products").Select("updated_at, code, price as prices").Scan(&out)
  // db.Table("products").Select("price as prices").Scan(&out) //{ 1000 }
  db.Table("products").Scan(&out) // { 0 L1212}
  fmt.Println(out)

}