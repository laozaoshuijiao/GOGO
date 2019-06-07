package main

import "fmt"

func MAXAndMIX(a , b int )(max ,min int){
	if a > b{
		max = a
		min = b
	} else{
		max = b
		min = a
	}
	return
}
func main(){
	max , min := MAXAndMIX(10,20)

	fmt.Printf("max = %d , min = %d\n",max , min)

	a , _ := MAXAndMIX(10,20)
	//通过mi名变量丢弃某个返回值
	fmt.Printf("a = %d\n",a)
}