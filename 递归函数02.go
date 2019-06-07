/*
许多问题都可以使用递归来解决，比如著名的快速排序算法
在使用递归函数是经常会遇到的一个重要问题就是宅一处，一般出现在
大量递归调用导致的程序扎内存分配耗尽，这个问题可以通过一个名为
懒惰求值的技术解决，在GO中，我们可以使用管道 
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
    fmt.Printf("%d is odd: is %t\n", 17, odd(17))
    // 17 is odd: is true
    fmt.Printf("%d is odd: is %t\n", 18, odd(18))
    // 18 is odd: is false
}

func even(nr int) bool {
    if nr == 0 {
        return true
    }
    return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
    if nr == 0 {
        return false
    }
    return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
    if nr < 0 {
        return -nr
    }
    return nr
}