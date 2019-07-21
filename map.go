/*
映射(Map): 将唯一键映射到值 

必须使用make函数来创建映射
// declare a variable, by default map will be nil
var map_variable map[key_data_type]value_data_type

// define the map as nil map can not be assigned any value
map_variable = make(map[key_data_type]value_data_type)
*/
package main

import "fmt"

func main() {
	var countryCapitalMap map[string] string

	countryCapitalMap = make(map[string] string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"

	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not*/
	// country, capital := countryCapitalMap["France"]
	// fmt.Println(country, capital)

	delete(countryCapitalMap, "France")

	for country, capital := range countryCapitalMap {
		fmt.Println(country, capital)
	}
}