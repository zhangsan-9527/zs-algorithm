package main

import "fmt"

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n //凡是递归，一定要有终止条件，否则会进入无限循环
	}
	return Fibonacci(n-1) + Fibonacci(n-2) //递归调用函数自身
}

// 1. 写出递归公式: f(n) = f(n-1) + f(n-2)
// 2. 写出递归的终止条件: f(1) = 1; f(2) = 2
func walkStair(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	return walkStair(n-1) + walkStair(n-2)
}

func main() {
	//fmt.Println(Fibonacci(2))
	fmt.Println(walkStair(5))
}
