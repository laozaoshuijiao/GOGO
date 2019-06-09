package main

import "fmt"

var a byte//全局变量

func main(){
	var a int//局部变量
	//不同作用域，允许同名变量
	//使用变量的原则，就近原则
	fmt.Printf("%T\n",a)
	{
		var a float32

		fmt.Printf("%T\n",a)
	}
	test()

}
func test(){
	fmt.Printf("%T\n",a)
}