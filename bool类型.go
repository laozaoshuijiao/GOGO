package main

import "fmt"

func main(){
	var a bool
	fmt.Println("a0 = ",a)
	a = true
	fmt.Println("a = ",a)

	//自动推导类型
	var b = false
	fmt.Println("b = ",b)

	c := false
	fmt.Println("c = ",c)
}
