// 4.5 define a alias named Rope for string, and declare a variable
// 4.5 定义一个 string 的类型别名 Rope，并声明一个该类型的变量
package main

import "fmt"

type Rope string

func main() {
	var str Rope = "Hey"
	fmt.Println(str)
}
