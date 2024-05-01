// 4.6 stats byte & rune for 2 strings. Hint: unicode/utf8 pkg
// 4.6 创建一个用于统计字节和字符（rune）的程序，对以下 2 个字符串进行分析（提示：使用 unicode/utf8 包）
// asSASA ddd dsjkdsjs dk
// asSASA ddd dsjkdsjsこん dk

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "asSASA ddd dsjkdsjs dk"
	s2 := "asSASA ddd dsjkdsjsこん dk"

	fmt.Println("Byte count of s1: ", len(s1))
	fmt.Println("Rune count of s1: ", utf8.RuneCountInString(s1))
	fmt.Println("Byte count of s2: ", len(s2))
	fmt.Println("Rune count of s2: ", utf8.RuneCountInString(s2))
}
