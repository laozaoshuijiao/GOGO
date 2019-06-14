package main

import "fmt"


func main(){
	//自动推导类型，同时初始化
	s1 := []int{1,2,3,4}
	fmt.Println("s1 = ",s1)
	//借助make函数make(类型，len，cap)
	s2 := make([]int,5,10)
	fmt.Println("s2 = ",s2)

	//没有指定容量,容量和长度一样
	s3 := make([]int ,10)
	fmt.Println("s3 = ",s3)
}

func main01(){
	//数组方括号里面的是一个固定的常量，数组不能修改长度
	a :=[5]int{}
	fmt.Printf("len = %d , cap = %d\n",len(a),cap(a))

	//切片，【】里面为空，或者为。。。切片里面的容量不固定
	s:=[]int{}
	fmt.Printf("len = %d , cap = %d\n",len(s),cap(s))
	s = append(s,11)
	fmt.Printf("append:len = %d , cap = %d\n",len(a),cap(a))
}
