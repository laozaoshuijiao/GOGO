package main

import "fmt"

//定义接口类型
type Humaner interface {
	//方法只有声明，没有实现,有别的类型(自定义类型)实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

//Student 实现了此方法
func (tmp *Student) sayhi(){
	fmt.Printf("Student[%s,%d] sayhi\n",tmp.name,tmp.id)
}

func main(){

}
