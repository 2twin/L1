package main

import "fmt"

func main() {
	a := 1
	b := 2

	// 1 вариант
	a, b = b, a
	fmt.Printf("a: %d\nb: %d\n", a, b)

	fmt.Println("---------")

	// 2 вариант
	a += b
	b = a - b
	a -= b
	fmt.Printf("a: %d\nb: %d\n", a, b)
}