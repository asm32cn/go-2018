// go-map-1.go
package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Idaly"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 查看元素在集合中是否存在 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["United States"]

	if(ok){
		fmt.Println("Capital of United States is ", capital)
	}else{
		fmt.Println("Capital of United States is not present")
	}
}