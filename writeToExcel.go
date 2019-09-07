package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
)

type Student struct {
    Name string
    Age int
    School string 
    HomeTown string
}

func main() {
    var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet1")
    if err != nil {
        fmt.Printf(err.Error())
    }

    // 设置单元格的样式
    cStyle  := xlsx.NewStyle()
    cStyle.Alignment.Horizontal = "center"

    col := sheet.Col(0)
    col.Width = 20
    col.SetStyle(cStyle)
    col = sheet.Col(1)
    col.Width = 10
    col.SetStyle(cStyle)
    col = sheet.Col(2)
    col.Width = 20
    col.SetStyle(cStyle)
    col = sheet.Col(3)
    col.Width = 20

    row = sheet.AddRow()
    cell = row.AddCell()
    cell.Merge(3, 1)
    cell.SetString("学生个人信息表")
    fStyle := xlsx.NewStyle()
    fStyle.Alignment.Horizontal = "center"
    fStyle.Alignment.Vertical = "center"
    fStyle.Font.Size = 15
    cell .SetStyle(fStyle)

    second := sheet.Row(2)
    cell = second.AddCell()
    cell.SetString("姓名")
    cell = second.AddCell()
    cell.SetString("年龄")
    cell = second.AddCell()
    cell.SetString("学校")
    cell = second.AddCell()
    cell.SetString("住所")

    datas := []*Student{}
    stu1 := &Student{"clarence", 20, "school1", "home1"}
    stu2 := &Student{"john", 10,  "school2", "home2"}
    datas = append(datas, stu1)
    datas = append(datas, stu2)

    for _, data := range datas{
        row := sheet.AddRow()
        row.WriteStruct(data, 4)
    }

    err = file.Save("学生个人信息表.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
}