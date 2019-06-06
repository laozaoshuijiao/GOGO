package main

import ("fmt"
		"time"	)


func main(){
	for i := 0; i < 100;i++{
		time.Sleep(time.Second)
		if i == 10{
			//break //跳出循环，如果嵌套多个循环，跳出最近的内循环
			continue//跳出本次循环，下一次继续
		}
		fmt.Println("i = ",i)
	}
}
