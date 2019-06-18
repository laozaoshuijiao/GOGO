/*
Go语言支持并发，我们只需要通过go关键字来开启goroutine即可
goroutine是轻量级线程，goroutine的调度是由golang运行时进行管理的
goroutine语法格式

go 函数名(参数列表)
例如：

go f (x , y ,z)

开启一个新的 goroutine
*/
/*GO允许使用go语句开启一个新的运行期线程，即goroutine，以一个不同的，新
创建的goroutine来执行一个函数。同一个程序中的所有goroutine共享同一个地址空间*/
package main

import (
	"fmt"
	"time"
)
func say(s string){
	for i := 0;i < 5 ; i ++{
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}
func main(){
	go say("world!")
	say("hello")
}
