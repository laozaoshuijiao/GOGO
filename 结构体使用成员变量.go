package main

import"fmt"

type Student struct{
	id int
	name string
	sex string
	age int
	addr string
}

func main(){
	//定义一个结构体普通变量
	var s Student

	//操作成员，需要使用运算符点（.)
	s.id = 1
	s. name = "mika"
	s.sex = "男"
	s.age = 18
	s.addr = "bj"

	fmt.Println("s = ",s )
}
