// 5.7 print 1 ~ 100
// if multiple of 3, print Fizz; if multiple of 5, print Buzz; if both, print FizzBuzz
// Hint: switch

// 5.7 写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。
// 遇到 5 的倍数时，打印 Buzz 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 FizzBuzz（提示：使用 switch 语句）

package main

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0:
			println("Fizz")
		case i%5 == 0:
			println("Buzz")
		case i%3 == 0 && i%5 == 0:
			println("FizzBuzz")
		default:
			println(i)
		}
	}
}
