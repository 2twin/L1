package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	// разбиваем строку на массив элементов, отделенных пробелом
	list := strings.Split(str, " ")
	// создаем два указателя (наименьший и наибольший индексы массива)
	left, right := 0, len(list) - 1

	// пока левый указатель меньше правого, меняем элементы массива по этим индексам местами
	for left < right {
		list[left], list[right] = list[right], list[left]

		// сдвигаем правый указатель на 1 влево, левый - на 1 вправо
		right--
		left++
	}

	// соединяем массив обратно в строку через пробел
	return strings.Join(list, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}