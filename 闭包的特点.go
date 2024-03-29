package main

import "fmt"

//函数的返回值是一个匿名函数，返回一个函数类型
func test02()func()int{
	var x int

	return func()int{
		x ++
		return x * x
	}
}

func main(){
	//返回值为一个匿名函数，返回一个函数类型,通过F来调用返回的匿名函数，f来调用闭包的函数
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func test01()int{
	//函数被调用时，x才分配空间，才初始化为0
	var x int //没有初始化值为0
	x ++
	return x * x//函数调用完毕，x自动释放
}

func main01(){
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())

}
