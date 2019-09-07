package main

import (
	"fmt"
	"errors"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "学生个人信息表.xlsx"
	var errNotFound error = errors.New("Can't find the file: " + excelFileName)

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("error: %v", errNotFound)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s ", text)
			}
			fmt.Println()
		}
	}
}
alter table boss_reserve_order modify column check_in_time datetime null not default "0000-00-00 00:00:00"