package main

import "fmt"

type Person struct {
	name string     // 名字
	sex  string     // 性别
	age  int        //年龄
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id int
	addr string
	name string //和Peckag同名
}

func main() {
	//声明定义一个变量
	var s Student

	//默认规则，如果能在本作用域找到此成员，就操作此成员，如果没找到就找到继承的字段
	s.name = "mike"
	s.sex = "m"
	s.age = 18
	s.addr = "bj"

	fmt.Printf("s = %+v\n",s)

	//显示调用
	s.Person.name = "yoyo"
	fmt.Printf("s = %+v\n",s)
}