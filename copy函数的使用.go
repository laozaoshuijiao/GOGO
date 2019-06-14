package main

import "fmt"

func main(){
	src := []int{1,2}
	src1 := []int{6,6,6,6,6}

	copy(src1,src)
	fmt.Println("src= ",src)
	fmt.Println("src1= ",src1)

}
