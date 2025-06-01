package main

import "fmt"

func main() {
	a := 5
	// var pointerA *int - это указатель на переменную int
	pointerA := &a // Создание указателя на переменную a
	// Оператор * dereference: значение по адресу памяти
	fmt.Println(*pointerA) // 5
	double(&a)
	fmt.Println(a) // 10
}

// Функция double принимает ссылку на int и ничего не возвращает
// В результате происходит мутирование (изменение) исходного объекта
func double(num *int) {
	// Мутирование исходного объекта
	*num = *num * 2 // *num - * получить значение из адреса в памяти
}
