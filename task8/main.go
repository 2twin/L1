package main

import "fmt"

func main() {
	// создаем переменную и получаем ее значение через stdin
	var num int64
	fmt.Scan(&num)

	fmt.Printf("Исходная переменная: %d - %b\n", num, num)

	// получаем номер бита, который нужно заменить через stdin
	var numOfBit int64
	fmt.Scan(&numOfBit )

	fmt.Printf("Заменить %d бит\n", numOfBit)

	// создаем "маску"
	var mask int64 = 1<<(numOfBit-1)

	// применяем XOR для замены нужного бита
	result := num ^ mask

	fmt.Printf("Полученная переменная: %d - %b\n", result, result)
}
 