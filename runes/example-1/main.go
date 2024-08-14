package main

import (
	"fmt"
	"unicode/utf8"
)

// Go 字符串是字节的只读切片。语言和标准库对字符串有特殊处理 - 将其作为 UTF-8 编码文本的容器。
// 在其他语言中，字符串由“字符”组成。在 Go 中，字符的概念称为符文 - 它是一个表示 Unicode 代码点的整数。

func main() {
	// s 是一个字符串，被赋予了代表泰语单词“hello”的字面值。Go 字符串字面值是 UTF-8 编码的文本。
	const s = "我是二狗"
	// 由于字符串相当于 []byte，因此这将产生存储在其中的原始字节的长度。
	fmt.Println(len(s))

	// 对字符串进行索引会生成每个索引处的原始字节值。此循环会生成构成 s 中代码点的所有字节的十六进制值。
	for i :=0; i< len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	// 要计算字符串中有多少个符文，我们可以使用 utf8 包。请注意，RuneCountInString 的运行时间取决于字符串的大小
	// 因为它必须按顺序解码每个 UTF-8 符文。一些泰语字符由可以跨越多个字节的 UTF-8 代码点表示，因此此计数的结果可能会令人惊讶
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	
}