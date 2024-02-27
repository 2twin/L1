package main

import "fmt"

type Sedan struct {
	Brand string
	Year  int
}

func (s Sedan) String() string {
	return fmt.Sprintf("Sedan { brand: %s, year of manufacture: %d }", s.Brand, s.Year)
}

// Встраиваем структуру Sedan в структуру Car
type Car struct {
	Sedan
	Mileage int
}

func (c Car) String() string {
	return fmt.Sprintf("Car { brand: %s, year of manufacture: %d, mileage: %d }", c.Brand, c.Year, c.Mileage)
}

func main() {
	car := Car{
		Sedan: Sedan{
			Brand: "Toyota",
			Year:  2006,
		},
		Mileage: 40_000,
	}

	// Видим, что у car есть поля Brand и Year встроенной структуры Sedan
	fmt.Println(car)
}
