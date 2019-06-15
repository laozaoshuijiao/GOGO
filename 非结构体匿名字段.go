package main

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string     // 名字
	sex  string     // 性别
	age  int        //年龄
}

type Student struct {
	Person //匿名字段，
	int   //基础类型的匿名字段
	addr string
	name string //和Peckag同名
	mystr

	func main(){
	s := Student{Person{"mike","m",18},666,"heheeheh"}
	fmt.
	}