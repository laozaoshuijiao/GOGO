package main

import "fmt"

func main(){
	// for 初始化条件； 判断条件 ； 条件变化{
	// }
	//1 + 2 + 3 + 。。。。。1000
	sum := 0
	//初始化条件年 i：= 1
	//判断条件是否为真 i《 100，如果为真就执行循环体，如果为假就跳出循环体
	//条件变化 i++
	//不停重复 判断条件 和条件变化
	for i := 1 ;i < 100 ; i ++{
		sum = sum + i
	}
	fmt.Println("sum = ",sum)
}
