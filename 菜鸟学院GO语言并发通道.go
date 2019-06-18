/*
通道是用来传递数据的一个数据结构
通道可用于两个goroutine之间通过传递一个指定类型的值来同步运行和通讯。操作符
<-用于指定通道方向，发送或接收。如果未指定

ch <- v //把 v 发送到通道ch
v := <-ch //从ch接收数据，并把值赋给v
声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须创建：
ch := make(chan int)

默认情况下，通道是不带缓冲区的，发送端发送数据，同时必须又接收端相应解释偶数据。

*/
//实例通过两个goroutine来计算数字之和，在goroutine完成计算后，他们会计算两个
//结果的和
package main

import"fmt"

func sum(s[]int,c chan int){
	sum := 0
	for _,v := range s{
		sum += v

	}
	c <- sum //把sum发送到通道c
}
func main(){
	s := []int{7,2,8,-9,4,0}

	c := make(chan int)

	go sum(s[:len(s)/2],c)
	go sum (s[len(s)/2:],c)

	x , y := <-c,<-c //从通道c 中接收

	fmt.Println(x,y,x+y)

}