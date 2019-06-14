package main

import (

	"fmt"
)
//数组做函数参数，它是值传递
//是参数组的每个元素给形参数组拷贝一份
//形参数组是这个实参数组的复制品
func modify(a [5] int){
	a[0] = 6666
	fmt.Println("modify = ",a)
}

func main(){
	a := [5]int{1,2,3,4,5} //初始化

	modify(a)//数组传递国区
	fmt.Println("main a =",a)
}