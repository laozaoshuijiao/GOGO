package main

import (
	"fmt"
	"os"
)

func main(){
	//接收用户传递的参数，都是以字符串方式传递的
	list := os.Args

	n := len(list)
	fmt.Println("n = ",n)

	for i := 0 ; i < n; i ++{
		fmt.Printf("list[%d] = %s\n",i,list[i])
	}
}