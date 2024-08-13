package main

import "fmt"

// 可变参数函数可以使用任意数量的尾随参数来调用。例如，fmt.Println 是一个常见的可变参数函数。

// 这是一个将任意数量的整数作为参数的函数。
// 在函数内部，nums的类型相当于 []int。我们可以调用 len(nums)、使用 range 对其进行迭代等。
func sum(nums ...int)  int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3))
	fmt.Println(sum(1,2,3,4))
}