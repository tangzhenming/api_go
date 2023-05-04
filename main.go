package main

import "fmt"

func main() {
	// Hello World
	println("Hello World")
	fmt.Println("Hello World in fmt")

	// Variable
	var a = 1
	var b int = 2
	var c, d = 3, 4
	var (
		e = 5
		f = 6
	)
	fmt.Println(a, b, c, d, e, f)
	// 一般都使用冒号等于号语法声明变量（无法声明类型但是会自动推断） :=
	g := 7
	h, i := 8, 9
	fmt.Println(g, h, i)

	// Constant
	const A = "A"
	const (
		B = iota
		C
		D
	)
	fmt.Println(A, B, C, D) // C D 从 B 开始自动推断（B 的值 iota 为希腊语的第九个字母，表示极微小；一般用来做枚举值，只用写第一个就行；只能用在 const 中）
	const (
		A1 = iota + 1
		_
		A2
		A3
	)
	fmt.Println(A1, A2, A3) // _ 表示跳过
	const (
		B1 = 1 << iota
		B2
		B3
		B4
	)
	fmt.Println(B1, B2, B3, B4) // 左移操作
}

// go build . && ./api_go
// or
// go run .
