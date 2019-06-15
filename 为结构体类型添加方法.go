package main

import"fmt"

type Person struct {
	name string
	sex  string
	age  int
}

//带有接收者的函数叫方法
func (tmp Person) Printlnfo(){
	fmt.Println("tmp = ", tmp)
}
//通过一个函数,给成员赋值
func (p *Person) SetInfo(n ,s string,a int){
	n = p.name
	s = p.sex
	a = p.age
}
func main(){
	//定义同时初始化
	s := Person{"miki","男",18}
	s.Printlnfo()

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo","女",22)
	p2.Printlnfo()
}