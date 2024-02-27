package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	lock  sync.RWMutex
}

// sumOfSquares конкурентно рассчитывает сумму квадратов всех элементов слайса list и выводит результат в stdout.
func sumOfSquares(list []int) {
	c := Counter{
		value: 0,
	}

	wg := sync.WaitGroup{}

	wg.Add(len(list))

	for _, val := range list {
		go func(val int) {
			defer wg.Done()

			c.lock.RLock()
			defer c.lock.RUnlock()

			c.value += val * val
		}(val)
	}

	wg.Wait()

	fmt.Println(c.value)
}

func main() {
	list := []int{2, 4, 6, 8, 10}
	sumOfSquares(list)
}
