package main

import "fmt"

func main(){
	a , b  := 10 , 20

	//交换两个变量的值
	var tmp int
	tmp = a
	a = b
	b = tmp

	fmt.Printf("a = %d ,b = %d\n",a ,b)

	i , j := 10 , 20
	i , j = j , i
	fmt.Printf("i = %d ,j = %d\n",i ,j)
	//迷悯变量，丢弃数据不处理
	tmp , _ = i , j
	fmt.Printf("tmp = %d \n",tmp)

}
