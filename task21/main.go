package main

import "fmt"

// создаем интерфейс Transport с методом Drive
type Transport interface {
	Drive()
}

// создаем тип Auto
type Auto struct{}

// имплементируем интерфейс Transport для Auto
func (a *Auto) Drive() {
	fmt.Println("Driving Auto")
}

// создаем тип Driver
type Driver struct {}

// создаем метод Travel для водителя, который принимает аргумент типа Transport и вызывает его метод Drive
func (d *Driver) Travel(t Transport) {
	t.Drive()
}

// создаем интерфейс Animal с методом Move
type Animal interface {
	Move()
}

// создаем тип Horse
type Horse struct{}

// имплементируем интерфейс Animal для Horse
func (h *Horse) Move() {
	fmt.Println("Riding Horse")
}

// создаем адаптер для типа Auto, который содержит поле типа Horse
type HorseToAutoAdapter struct {
	horse *Horse
}

// имплементируем интерфейс Auto для адаптера
func (adapter *HorseToAutoAdapter) Drive() {
	adapter.horse.Move()
}

func main() {
	// создаем водителя, автомобиль и отправляем водителя в путешествие на авто
	driver := &Driver{}
	auto := &Auto{}
	driver.Travel(auto)

	// теперь водителю нужно пересесть на лошадь: создаем переменную типа Horse и адаптер для Auto с полем horse
	horse := &Horse{}
	horseTransoprt := &HorseToAutoAdapter{
		horse: horse,
	}

	// водитель теперь может спокойно передвигаться на лошади посредством адаптера
	driver.Travel(horseTransoprt)
}