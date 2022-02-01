package main

import "fmt"

func main() {
	fmt.Println(soma(10, 5))
}

func soma(a int, b int) int {
	return a + b
}

func multi(a int, b int) int {
	return a * b
}

func sub(a int, b int) int {
	return a - b
}
