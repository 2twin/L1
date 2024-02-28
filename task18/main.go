package main

import (
	"fmt"
	"sync"
)

// создаем структуру с мьютексом, которая хранит значение счетчика
type Counter struct {
	value int
	lock  sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
	}
}

func (c *Counter) Increment() {
	// блокируем счетчик и увеличиваем значение на 1
	c.lock.Lock()
	defer c.lock.Unlock()

	c.value++
}

func (c *Counter) GetValue() int {
	// блокируем счетчик на чтение и возвращаем его значение
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.value
}


func main() {
	c := NewCounter()

	// создаем wait группу
	wg := sync.WaitGroup{}
	wg.Add(50)

	// запускаем 50 горутин, в каждой из которых увеличиваем значение счетчика на 1
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	// дожидаемся выполнения всех горутин и выводим значение счетчика
	wg.Wait()
	fmt.Println(c.GetValue())
}