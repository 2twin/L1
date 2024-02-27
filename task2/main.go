package main

import (
	"fmt"
	"sync"
)

// square конкурентно возводит в квадрат каждый элемент слайса list и выводит их в stdout.
func square(list []int) {
	wg := sync.WaitGroup{}
	wg.Add(len(list))

	for _, val := range list {
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(val)
	}

	wg.Wait()
}

func main() {
	list := []int{2, 4, 6, 8, 10}
	square(list)
}
