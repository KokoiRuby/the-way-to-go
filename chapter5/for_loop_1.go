// 5.4.1 Create a simple for loop, print counter 15 times using fmt package
// 5.4.1 使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值

// 5.4.2 Using goto, do not use for keyword
// 5.4.2 使用 goto 语句重写循环，要求不能使用 for 关键字

package main

import "fmt"

func main() {
	var counter int
	/*	for i := 0; i < 15; i++ {
		fmt.Println(counter)
		counter++
	}*/
ORIGIN:
	fmt.Println(counter)
	counter++
	if counter == 15 {
		return
	}
	goto ORIGIN
}
