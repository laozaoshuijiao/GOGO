package main

import "fmt"

type Student struct{
	id int
	name string
	sex string
	age int
	addr string
}

func main(){
	//指针有合法对象后，才操作成员
	//先定义一个普通结构体变量
	var s Student
	//定义一个指针变量保存 s 的地址
	var p1 *Student

	p1 = &s

	//操作成员 p1.id 和(*p1)。id万全等价,使用.运算符
	p1.id = 18
	(*p1).name = "make"
	p1.age = 18
	p1.sex = "m"
	p1.addr = "bj"
	fmt.Println("p1 = ",p1)

	//同
	p2 := new(Student)
	p2.id = 18
	p2.name = "make"
	p2.age = 18
	p2.sex = "m"
	p2.addr = "bj"
	fmt.Println("p2 = ",p2)
}
