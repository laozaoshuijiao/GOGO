package main

import "fmt"

func main(){
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaa")
	//defer延迟调用，main函数结束前调用
	defer fmt.Println("bbbbbbbbbbbbbbbbbbbbbbb")
}
