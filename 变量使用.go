package main

import "fmt"
func main(){
	//1.声明 格式 var 变量名 类型
	var a int //变量声明了必须使用

	fmt.Println(a)

	a = 10 //变量赋值
	fmt.Println(a)

	var b int = 10  //初始化
	fmt.Println(b)
	b = 20          //赋值
	fmt.Println(b)

	//自动推导类型，必须初始化，通过初始化的值确定类型
	c := 30
	//%Td
	fmt.Print("c type is %T\n",c)




}
