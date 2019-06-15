package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}
func test01(s Student){
	s.id = 666
	fmt.Println("test01 :",s)
}

func test02(p* Student){
	p.id = 666
	fmt.Println("test02 :",p)
}

func main(){
	s :=Student{1,"mike","男",18,"bj"}

	test02(&s) //值传递,形参无法改实参，地址传递 形参可以改实参
	fmt.Println("main :",s)
}
