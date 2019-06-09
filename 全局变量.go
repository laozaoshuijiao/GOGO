package main

import "fmt"

//定义在函数外部的变量叫做全局变量
//全局变量在全局都可以使用

func test(){
	fmt.Println("test a = ",a)
}

var a int

func main(){
	a = 10
	fmt.Println("a = ",a)
	test()
}

