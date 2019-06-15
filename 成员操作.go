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
}

func main() {
	//顺序初始化
	s1  := Student{Person{"miki", "男", 18}, 1, "bj"}
	s1.Person =Person{age : 17}
	s1.name = "yoyo"
	fmt.Println(s1.name,s1.age,s1.id)
}
