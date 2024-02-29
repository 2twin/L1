package main

import (
	"errors"
	"fmt"
)

func delete(list *[]int, idx int) error {
	// идем по указателю и достаем сам массив
	l := *list

	// елси индекс выходит за рамки массива, возвращаем ошибку
	if idx < 0 || idx >= len(l) {
		return errors.New("index out of range")
	}
	
	// модифицируем исходный массив, убрав из него элемент по индексу idx
	*list = append(l[:idx], l[idx+1:]...)
	return nil
}

func main() {
	list := []int{1, 23, 4, 56, 7, 89}
	
	err := delete(&list, 6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)

	err = delete(&list, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}
