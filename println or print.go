package main

import "fmt"

func main(){
	a := 10
	//一段一段处理，自动换行输出
	fmt.Println("a = ",a)
	//格式化输出，把a的内容放在%d的位置

	fmt.Printf("a = %d\n",a)

	b , c := 20 , 30
	fmt.Printf("a = %d b = %d ,c = %d",a,b,c)
}
