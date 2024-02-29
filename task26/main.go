package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	// создаем мапу для отслеживания уникальных символов
	m := make(map[rune]struct{})

	// приводим строку к нижнему регистру
	lower := strings.ToLower(str)

	// итерируемся по строке, и если в мапе уже есть такой ключ, возвращаем false
	for _, ch := range lower {
		if _, ok := m[ch]; ok {
			return false
		}

		// в противном случае, записываем руну в мапу
		m[ch] = struct{}{}
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(isUnique(s1))
	fmt.Println(isUnique(s2))
	fmt.Println(isUnique(s3))
}