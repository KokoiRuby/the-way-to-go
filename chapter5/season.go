// 5.2 Write a Season func, accept a number that represents month, and return corresponding season
// 5.2 写一个 Season 函数，要求接受一个代表月份的数字，然后返回所代表月份所在季节的名称（不用考虑月份的日期）
package main

import (
	"fmt"
)

func main() {
	var month int
	fmt.Println("Please input a month: ")
	input, err := fmt.Scanln(&month)
	if err != nil {
		return
	}
	Season(input)
}

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Winter"
	default:
		return "Illegal Month!"
	}
}
