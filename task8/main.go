package main

import "fmt"

func main() {
	var num int64
	fmt.Scan(&num)

	fmt.Printf("Исходная переменная: %d - %b\n", num, num)

	var numOfBit int64
	fmt.Scan(&numOfBit )

	fmt.Printf("Заменить %d бит\n", numOfBit)

	var mask int64 = 1<<(numOfBit-1)
	result := num ^ mask

	fmt.Printf("Полученная переменная: %d - %b\n", result, result)
}
 