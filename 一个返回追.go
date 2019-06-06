package main

import "fmt"

//无参有返回值：只有一个返回值
//有返回值的函数需要通过return中断返回
func myfunc01() int {
	return 666
}
//给返回值取一个变量名,go语言推荐写法
func myfunc02() (result int) {
	return 666
}
//常用的一个写法
func myfunc03() (results int) {
	results = 666
	return
}

func main(){
	//无参有返回值函数调用
	var a int
	a = myfunc01()
	fmt.Println("a = ",a)

	b := myfunc01()
	fmt.Println("b = ",b)

	c := myfunc03()
	fmt.Println("c = ",c)
}