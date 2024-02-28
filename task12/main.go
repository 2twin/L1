package main

import "fmt"

// объявляем тип для нашего set
type set[T comparable] map[T]struct{}

// NewSet возвращает укакзатель на новый сет, созданный из переданного слайса.
func NewSet[T comparable](list []T) *set[T] {
	s := make(set[T])

	// итерируемся по каждому элементу слайса, если его нет в set - добавляем
	for _, val := range list {
		if _, ok := s[val]; !ok {
			s[val] = struct{}{}
		}
	}

	return &s
}

func main() {
	list := []string{"cat", "cat", "dog", "cat", "tree"}

	s := NewSet(list)

	fmt.Println(s)
}