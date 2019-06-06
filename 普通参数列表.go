package main

import "fmt"

//有参无返回值函数的定义,普通参数列表
//定义函数时，在函数名后面（）定义的叫形参
//参数传递，只能由实参传递给形参，不能反过来，单向传递
func Myfunc01(a int){
	//a = 1111
	fmt.Println("a = ",a)
}
func Myfunc02(a int , b int){
	fmt.Printf("a = %d , b = %d", a , b)
}
func Myfunc03(a , b int){

}
func Myfunc04(a int , b string , c float64){

}
func main(){
	//有参无返回值函数调用：函数名（所需参数）
	//调用函数传递的参数 叫实参
	Myfunc01(6666)
	Myfunc02(666,7777)
}
