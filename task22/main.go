package main

import (
	"fmt"
	"math/big"
)

func main() {
	// используем пакет big для работы с большими числами
	a := big.NewInt(1 << 62)
	b := big.NewInt(1 << 60)

	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Printf("a * b = %s\n", mul.String())

	div := new(big.Int)
	div.Div(a, b)
	fmt.Printf("a / b = %s\n", div.String())

	add := new(big.Int)
	add.Add(a, b)
	fmt.Printf("a + b = %s\n", add.String())

	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Printf("a - b = %s\n", sub.String())
}