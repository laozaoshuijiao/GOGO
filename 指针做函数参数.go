package main

import "fmt"

func swap(p1 , p2 *int){
	*p1 , *p2 = *p2, *p1

}

func main(){
	a, b := 10 , 20

	//通过一个函数交换a和b的内容
	swap(&a , &b) //地址传递
	fmt.Printf("a = %d b = %d\n",a,b)
}
