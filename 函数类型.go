package main

import "fmt"

func Add(a , b int)int{
	return a + b
}
func mainc(a , b int)int{
	return a - b
}
//函数也是一种数据类型，通过type给一个函数类型取名
//FuncType他是一个函数类型
type FuncType func (int,int)int //没有函数名字没有{}
func main(){
	var r int
	r = Add(1,1)//传统调用方式
	fmt.Printf("r = %d\n",r)

	//声明一个函数类型的变量，变量名叫做fTest
	var fTest FuncType

	fTest = Add //是变量就可以赋值
	r2 := fTest(10,20)//等价于Add（10 ， 20）
	fmt.Println(" r2= ",r2)

	fTest = mainc

	r3 := fTest(20 , 10)
	fmt.Println("r3 = ",r3)
}

