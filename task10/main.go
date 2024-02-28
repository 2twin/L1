package main

import "fmt"

func main() {
	// слайс с исходными данными
	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// создаем мапу, которая будет хранить все подмножества с шагом в 10
	m := make(map[int][]float64)

	for _, val := range list {
		// каждый элемент слайса аппендим в слайс под нужный ключ
		k := int(val) - int(val)%10
		m[k] = append(m[k], val)
	}

	fmt.Println(m)
}