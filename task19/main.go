package main

import "fmt"

func reverseString(str string) string {
	// создаем слайс рун из исходной строки
	res := []rune(str)

	// воспользуемся двумя указателями
	left := 0
	right := len(res) - 1

	// пока левый индекс меньше правого, меняем местами элементы в слайсе по этим индексам
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	return string(res)
}

func main() {
	str := "главрыба"
	rev := reverseString(str)
	fmt.Println(rev)
}