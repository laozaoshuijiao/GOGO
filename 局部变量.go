package main

import "fmt"

func test(){
	a := 10
	fmt.Println("a =",a)
}

func main(){
	//局部变量定义：定义在{}里面的变量就睡局部变量，只能里面
	//执行到定义变量那句话，才开始分配空间，离开时作用于自动释放
	//作用域，变量起作用范围
	{
		i := 10
		fmt.Println("i = ",i)
	}

}
