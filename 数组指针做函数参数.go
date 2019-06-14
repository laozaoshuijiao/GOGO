package main

import "fmt"

//p指向实现数组a，它指向数组，同时数组的指针
//*派代表的指向的内存，就是是实参a

func modify(p *[5]int){
	(*p)[0] = 666
	fmt.Printf("modify *a = %d\n",*p)
}

func main(){
	a := [5] int{1,2,3,4,5}

	modify(&a)//地址传递
	fmt.Println("a = ",a)

}

