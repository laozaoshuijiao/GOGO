/*
通道可以设置缓冲区，通过make的第二个参数指定缓冲区大小：
ch := make(chan , int , 100)
*/
package main

import"fmt"

func main(){
	 //这里我们定义了一个可以储存整数类型的带缓冲通道
	 //缓冲区顶通过为2
	 ch := make(chan  int ,3)

	 //因为 ch 是带缓冲通道，我们可以同时发送两个数据
	 //而不用立刻需要去同步数据
	 ch <- 1
	 ch <- 2
	 ch <- 3
	 //获取这两个数据
	 fmt.Println(<- ch)
	 fmt.Println(<- ch)
	 fmt.Println(<- ch)
}