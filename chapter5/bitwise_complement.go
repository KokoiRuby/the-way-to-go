// 5.6 print %b for bit complement from 0 - 10
// 5.6 使用按位补码从 0 到 10，使用位表达式 %b 来格式化输出
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%b\n", ^i)
	}
}
