package main

import"fmt"

type Person struct {
	name string
	sex  string
	age  int
}
type Student struct {
	*Person   //指针类型
	id    int
	addr  string
}



func main(){
	s1 := Student{&Person{"miki","男",18} , 666,"北京"}
	fmt.Println(s1.addr,s1.name,s1.id,s1.age,s1.sex)

	//先定义变量

	var s2 Student
	s2.Person = new(Person) //分配空间
	s2.name = "yoyo"
	s2.sex = "女です"
	s2.age = 23
	s2.id = 222
	s2.addr = "ss"
	fmt.Println( //分配空间
	s2.name ,
	s2.sex ,
	s2.age ,
	s2.id ,
	s2.addr)
}