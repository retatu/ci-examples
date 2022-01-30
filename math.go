package main

import "fmt"

func main() {
	fmt.Println("Teste")
	fmt.Println(soma(10, 5))
}

func soma(a int, b int) int {
	return a + b
}
