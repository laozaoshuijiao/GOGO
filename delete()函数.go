//delete()函数用于删除集合元素，参数为map和其对应的key
package main

import "fmt"

func main(){
	//创建map
	countryCapitalMap := map[string]string{"France" : "Paris" , "Italy": "Rome" ,
		"Japan" : "Tokyo" , "India" : "New delhi" }

	fmt.Println("原始地图")

	//打印地图
	for country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

	//删除元素
	delete(countryCapitalMap,"France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除后的地图")

	//打印地图
	for country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}
}