// 5.8 print a 20 x 10 rect made of *
// 5.8 使用 * 符号打印宽为 20，高为 10 的矩形

package main

func main() {
	const (
		length = 20
		width  = 10
	)
	for i := 1; i <= width; i++ {
		symbol := "*"
		for j := 1; j <= length; j++ {
			symbol += "*"
		}
		println(symbol)
	}
}
