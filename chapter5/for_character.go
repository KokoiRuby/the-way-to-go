// 5.5 Print like below until each row has 25 chars
// a. for x2
// b. for x1 + string concat

// 5.5 创建一个程序，要求能够打印类似下面的结果（直到每行 25 个字符时为止）
// a. 使用 2 层嵌套 for 循环
// b. 使用一层 for 循环以及字符串连接

// G
// GG
// GGG
// GGGG
// GGGGG
// GGGGGG
package main

func main() {
	// a
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// b
	content := "G"
	for i := 1; i <= 25; i++ {
		println(content)
		content += "G"
	}

}
