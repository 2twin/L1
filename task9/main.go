package main

import (
	"fmt"
)

// writeToChan создает канал out и записывает в него все значения из слайса, возвращает out.
func writeToChan(list []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, val := range list {
			out <- val
		}
	}()

	return out
}

// doubleValues создает канал out, читает данные из входного канала и записывает их в out, возвращает out.
func doubleValues(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for val := range in {
			out <- val * 2
		}
	}()

	return out
}

func main() {
	// создаем слайс
	list := []int{1, 2, 3, 4, 6}
	// делаем ковеер чисел
	in := writeToChan(list)
	out := doubleValues(in)
	// выводим результат в stdout
	for val := range out {
		fmt.Println(val)
	}
}