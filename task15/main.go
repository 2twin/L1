package main

var justString string

// Проблема 1: используется срез строки, что может привести к неожиданному результату, например, если в строке есть символы, содержащие более 1 байт.
// Проблема 2: неэффективное использование памяти - мы используем лишь первые 100 элементов, но глобальная переменная justString все равно останется в памяти. 
// Решение: создание slice рун из строки v
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}