package main

import "fmt"

//实现 1 + 2 +  3 +....100

func test01()(sum int){
	for i := 1 ; i <= 100 ; i ++{
		sum += i
	}
	return
}

func main(){
	var sum int
	sum = test01()
	fmt.Println("sum = ",sum)
}
