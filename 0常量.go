package main

import "fmt"

func main(){
	//变量，程序运行期间可以改变的量,变量声明需要var
	//常量，程序运行期间不可以改变的量，常量声明需要const

	const a int = 10
	fmt.Println("a = ",a)
	const b = 10.0  //没有使用：=
	fmt.Printf("b = %T",b)
 }