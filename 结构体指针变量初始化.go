package main

import "fmt"

type Student struct {
	id int
	name string
	sex string //字符类型
	age int
	addr string
}

func main(){
	//顺序初始化,每个成员都有必须初始化
	var p1 *Student = &Student{1,"mike","男",18,"bj"}
	fmt.Println("*p1 = ",*p1)

	//指定成员初始化，没有初始化的成员，没有初始化的成员自动赋值为0
	p2 := &Student{ id : 1, name : "mike"}
	fmt.Println("p2" , p2)
}
