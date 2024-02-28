package main

import "fmt"

// объявляем тип set
type set[T comparable] map[T]struct{}

func main() {
	// создаем 2 сета
	set1 := set[int]{
		2: {},
		1: {},
		4: {},
		3: {},
	}

	set2 := set[int]{
		4: {},
		2: {},
		6: {},
		8: {},
	}

	// создаем пустой set
	intersect := make(set[int])

	// итерируемся по всех элементам (ключам) set1, если такой элемент есть и в set2,
	// кладем его в intersect
	for k := range set1 {
		if _, ok := set2[k]; ok {
			intersect[k] = struct{}{}
		}
	}

	fmt.Println(intersect)
}
