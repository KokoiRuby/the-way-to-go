// 5.1 Modify string_conversion2.go, use := to assign value to err
// 5.1 尝试改写 string_conversion2.go 中的代码，要求使用 := 方法来对 err 进行赋值，哪些地方可以被修改？
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// assign error
	an, err := strconv.Atoi(orig)
	// return to exit if err
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
