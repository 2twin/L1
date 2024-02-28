package main

import "fmt"

func quicksort(list []int) {
	// если в слайсе 1 или меньше элементов, сортировать уже нечего
	if len(list) <= 1 {
		return
	}

	// задаем крайний левый и крайний правый индексы
	left, right := 0, len(list)-1

	// выбираем опорный элемент
	pivot := len(list) / 2

	// перемещаем опорный элемент в конец слайса
	list[pivot], list[right] = list[right], list[pivot]
	// не забываем свапнуть сами индексы
	pivot, right = right, pivot

	// итерируемся по слайсу
	// если i-ый элемент меньше опорного, то мы помещаем его на позицию left и увеличиваем значение left
	for i := 0; i < len(list); i++ {
		if list[i] < list[pivot] {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}

	// перемещаем опорный элемент после наименьшего значения
	list[left], list[pivot] = list[pivot], list[left]
	// не забываем свапнуть сами индексы
	left, pivot = pivot, left

	// сортируем части слайса, находящиеся слева и справа от опортного элемента
	quicksort(list[:pivot])
	quicksort(list[pivot+1:])

}

func main() {
	list := []int{1, 5, -4, 2, 12, -5, 0, 2}
	fmt.Println(list)

	quicksort(list)
	fmt.Println(list)
}
