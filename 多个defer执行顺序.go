package main


import"fmt"

func test(x int){
	r := 100 / x
	fmt.Println("r = ",r)
}

func main(){
	defer fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaa")


	defer fmt.Println("bbbbbbbbbbbbbbbbbbbbbbb")

	//调用一个函数导致内存出问题
	defer test(0)

	defer fmt.Println("ccccccccccccccccccccccccccccc")
}
