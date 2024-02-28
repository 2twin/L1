package main

import (
	"fmt"
	"reflect"
)

// проверка через type assertion
func typeof(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	// тут есть недостаток - проверять каждый тип канала вручную
	case chan int, chan bool, chan any:
		return "chan"
	default:
		return "invalid type"
	}
}

// альтернативный вариант - нет проблем с каналами
func typeOf(val interface{}) string {
	return reflect.TypeOf(val).Kind().String()
}

func main() {
	var val interface{} = 1
	fmt.Println(typeof(val))
	
	val = "abc"
	fmt.Println(typeof(val))
	
	val = true
	fmt.Println(typeof(val))
	
	val = make(chan int)
	fmt.Println(typeOf(val))
}