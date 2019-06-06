package main

import "fmt"

func test(tmp ...int) {
	for _,data :=range tmp{
		fmt.Println("tmp = ",data)
	}
}

func Myfunc(args ...int)  {
	//全部元素传递给Myfunc

	//
	Myfunc(args ...)
}
func main(){
	test(1,2,3,4)
}
