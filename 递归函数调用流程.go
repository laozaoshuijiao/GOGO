package main

import "fmt"

func test(a int){
	//函数调用自己本身
	test(a - 1)
	fmt.Println("a = ",a)
}

func main(){
	test(3)
	fmt.Println("main")
}
