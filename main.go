package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(fib(0))
	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}
