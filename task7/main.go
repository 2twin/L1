package main

import (
	"fmt"
	"sync"
)

// создаем структуру с мьютексом и мапой
type ConcMap struct {
	values map[int]string
	lock   sync.Mutex
}

func NewConcMap() *ConcMap {
	return &ConcMap{
		values: make(map[int]string),
	}
}

func (m *ConcMap) Set(k int, v string) {
	// блокируем мапу на запись и отложенно разблокируем
	m.lock.Lock()
	defer m.lock.Unlock()

	// записываем данные в мапу
	m.values[k] = v
}

func main() {
	concMap := NewConcMap()

	wg := sync.WaitGroup{}

	i := 15
	wg.Add(i)

	// запускаем 15 горутин, каждая из которых пишет данные в мапу
	for i > 0 {
		go func(i int) {
			defer wg.Done()
			concMap.Set(i, fmt.Sprintf("string #%d", i))
		}(i)
		i--
	}

	// дожидамеся выполнения всех горутин
	wg.Wait()

	fmt.Println(concMap)
}
