package main

import "fmt"

// go build . && ./api_go
// or
// go run .

// ======

// function
// go 中使用 func 声明函数时，函数内部不能再使用 func 嵌套声明函数，但是可以声明匿名函数（并赋值给变量）
// 还可以使用立即执行函数
func foo(a int, b int) int {
	// func bar () {} // expected expressionsyntax
	c := func() int {
		return 3
	}
	d := func(d1 int) int {
		return d1
	}(4)
	return a + b + c() + d
}

// go 中函数还可以返回多个值
// 返回值可以提前定义变量名字，并在 return 时默认返回
func bar(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

// 可变参数函数
func sum(numbers ...int) int {
	fmt.Println("numbers", numbers)
	total := 0
	for i, v := range numbers {
		fmt.Println(i, v) // 不想要 index 下标的话，可以使用 _ 代替 i
		total += v
	}
	return total
}

func main() {
	// // Hello World
	// println("Hello World")
	// fmt.Println("Hello World in fmt")

	// ======

	// // Variable
	// var a = 1
	// var b int = 2
	// var c, d = 3, 4
	// var (
	// 	e = 5
	// 	f = 6
	// )
	// fmt.Println(a, b, c, d, e, f)
	// // 一般都使用冒号等于号语法声明变量（无法声明类型但是会自动推断） :=
	// g := 7
	// h, i := 8, 9
	// fmt.Println(g, h, i)

	// ======

	// // Constant
	// const A = "A"
	// const (
	// 	B = iota
	// 	C
	// 	D
	// )
	// fmt.Println(A, B, C, D) // C D 从 B 开始自动推断（B 的值 iota 为希腊语的第九个字母，表示极微小；一般用来做枚举值，只用写第一个就行；只能用在 const 中）
	// const (
	// 	A1 = iota + 1
	// 	_
	// 	A2
	// 	A3
	// )
	// fmt.Println(A1, A2, A3) // _ 表示跳过
	// const (
	// 	B1 = 1 << iota
	// 	B2
	// 	B3
	// 	B4
	// )
	// fmt.Println(B1, B2, B3, B4) // 左移操作

	// ======

	// // for
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// j := 1
	// for {
	// 	if j > 3 {
	// 		break
	// 	}
	// 	fmt.Println("如果没有 break 就会无限循环", j)
	// 	j++
	// }

	// ======

	// // if
	// if num := 9; num < 0 { // 局部作用域，改变量仅在当前的 if else 中生效
	// 	fmt.Println(num, "负数", num)
	// } else if num < 10 {
	// 	fmt.Println("一位数", num)
	// } else {
	// 	fmt.Println("多位数", num)
	// }

	// num := 9
	// if num < 0 {
	// 	fmt.Println(num, "负数", num)
	// } else if num < 10 {
	// 	fmt.Println("一位数", num)
	// } else {
	// 	fmt.Println("多位数", num)
	// }

	// // switch
	// // switch 一个 case 可以跟多个值
	// // 默认无需使用 break
	// // 如果需要贯穿条件，可以加 fallthrough
	// switch time.Now().Weekday() {
	// case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
	// 	fmt.Println("It's weekday", time.Now().Weekday())
	// 	// fallthrough
	// default:
	// 	fmt.Println("It's weekend")
	// }

	// ======

	// function
	fmt.Println(foo(1, 2))

	fmt.Println(bar(1, 2))
	a, b := bar(3, 4)
	fmt.Println(a, b)

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
