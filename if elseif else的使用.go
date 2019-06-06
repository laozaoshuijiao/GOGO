package main

import "fmt"

func main(){
	if a := 11;a == 10{
		fmt.Println("a == 10")
	}else{//else没有条件
		fmt.Println("a != 10")
	}
	a := 10
	fmt.Println(a)
}
