package main

import "fmt"

func main(){
	array := []int{0,1,2,3,4,5,6,7,8,9}
	//取下标从low开始的元素，len = high - low,cap = max - low
	s1 :=array[:]
	fmt.Println("s1 =",s1)
	fmt.Printf("len = %d,cap = %d\n",len(s1),cap(s1))

	//操作某个元素，和数组操作方式一样
	data := array[0]
	fmt.Println("dara =",data)

	s2 := array[3:6:7]//a
	fmt.Println("s2 =",s2)
	fmt.Printf("len = %d,cap = %d\n",len(s2),cap(s2))

	s3 := array[:6]
	fmt.Println("s3 =",s3)
	fmt.Printf("len = %d,cap = %d\n",len(s3),cap(s3))
}
