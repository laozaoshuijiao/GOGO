package main

import "fmt"

func main(){
	m := map[int]string{1 : "mike" , 2 : "yoyo" , 3 : "go"}
	//第一个返回值为key，第二个返回值为value，编历结果是无序的
	for key,value := range m{
		fmt.Printf("%d =======> %v\n",key,value)
	}

	//如何判断一个key值是否存在
	//第一个返回值是返回key值所对应的value，第二个返回值为key是否存在的条件
	value, ok := m[0]

	 if ok == true {
	 	fmt.Println("m[1] = ",value )
	 }else{
	 	fmt.Println("不存在")
	 }



}
