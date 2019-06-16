package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}
//Person类型，实现一个方法
func (temp *Person)Printinfo(){
	fmt.Printf("name = %v","sex = %v，age = %v ",temp.name,temp.sex,temp.age)
}

//有一个学生,继承Person字段，成员方法都继承了
type Student struct {
	Person   //匿名字段
	id    int
	addr   string

}
func main(){
	s := Student{Person{"miki","男",18},666,"北京"}
	s.Printinfo()
}