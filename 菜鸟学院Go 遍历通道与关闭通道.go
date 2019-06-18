/*
GO通过range关键字 来实现编历读取到的数据，类似于与数组或切片。格式如下：
v,ok := <-ch

如果通道接收不到数据后ok就为false，这时通道就可以使用close（）函数来关闭
*/
package main

import "fmt"

func fibonacci (n int, c chan int){
	x , y := 0 , 1
	for i := 0 ; i < n ; i ++{
		c <- x
		x , y = y , x+y
	}
	close(c)
}
func main(){
	c := make(chan int , 10)
	go fibonacci(cap(c),c)
	//range 函数编历每个从通道接收的数据，因为c在发完10个
	//数据之后就关闭了通道，所以我们这range函数在接收10个数据
	//后就会关闭，如果上面的c通道不关闭，那么range函数就不会结束
	//从而在接收11个数据的时候就堵塞了。
	for i := range c{
		fmt.Println(i)
	}
}